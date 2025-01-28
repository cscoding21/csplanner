package project

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/marshal"
	"csserver/internal/services/organization"
	"csserver/internal/services/project/ptypes/projectstatus"
	"csserver/internal/services/resource"
	"fmt"

	"github.com/cscoding21/csval/validate"
)

var resourceMap map[string]resource.Resource

// SaveProject saves a project in the system and updates all stored calculations
func (s *ProjectService) SaveProject(ctx context.Context, pro Project, org organization.Organization) (common.UpdateResult[Project], error) {
	//---TODO: update to proper validation
	//val := pro.Validate()
	val := validate.NewSuccessValidationResult()

	if len(pro.ID) == 0 {
		return s.newProject(ctx, pro, org, val)
	}

	lastProject, err := s.GetProjectByID(ctx, pro.ID)
	if err != nil {
		//---existing project does not exist...create
		return s.newProject(ctx, pro, org, val)
	}

	//---ensure that the project status is not set through a back-channel
	pro.ProjectBasics.Status = lastProject.ProjectBasics.Status

	lastProject.ProjectBasics = pro.ProjectBasics
	lastProject.ProjectValue = pro.ProjectValue
	lastProject.ProjectCost = pro.ProjectCost
	lastProject.ProjectDaci = pro.ProjectDaci

	lastProject.GetProjectInitialCost()
	lastProject.GetProjectNPV()
	lastProject.GetProjectIRR()
	s.CalculateProjectMilestoneStats(lastProject)

	rm, err := s.GetResourceMap(false)
	if err != nil {
		return common.UpdateResult[Project]{Object: lastProject}, err
	}

	lastProject.CalculateProjectTaskStats(org, rm)

	return s.UpdateProject(ctx, lastProject)
}

func (s *ProjectService) newProject(ctx context.Context, pro Project, org organization.Organization, val validate.ValidationResult) (common.UpdateResult[Project], error) {
	pro.ProjectBasics.Status = projectstatus.NewProject
	proj, err := s.CreateProject(ctx, &pro)
	if err != nil {
		return common.NewUpdateResult(&val, &pro), err
	}

	pro = *proj.Object

	pro.GetProjectInitialCost()
	pro.GetProjectNPV()
	pro.GetProjectIRR()
	s.CalculateProjectMilestoneStats(&pro)

	rm, _ := s.GetResourceMap(false)

	pro.CalculateProjectTaskStats(org, rm)

	return s.UpdateProject(ctx, &pro)
}

// SetProjectStatus update the status of a project by adhering to the rules of the state machine
func (s *ProjectService) SetProjectStatus(ctx context.Context, projectID string, status projectstatus.ProjectState) ([]StatusCondition, error) {
	p, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return []StatusCondition{}, err
	}

	newState := stateMachineMap[status]

	conditions, err := newState.Can(p)
	if err != nil {
		return conditions, err
	}

	can := true
	for _, c := range conditions {
		if !c.Met {
			can = false
		}
	}

	if !can {
		return conditions, fmt.Errorf("cannot update project state...conditions not met")
	}

	p.ProjectBasics.Status = newState.State
	_, err = s.UpdateProject(ctx, p)

	return conditions, err
}

// GetResourceMap return a map of all resources keyed by ID
func (s *ProjectService) GetResourceMap(force bool) (map[string]resource.Resource, error) {
	if resourceMap != nil && force {
		return resourceMap, nil
	}

	p := common.NewPagedResultsForAllRecords[resource.Resource]()
	sql := "SELECT * from resource"
	res, resultCount, err := s.DBClient.FindPagedObjects(sql, p.Pagination, p.Filters)
	if err != nil {
		return nil, err
	}

	m := make(map[string]resource.Resource)
	p.Pagination.TotalResults = &resultCount
	unpacked, err := marshal.SurrealSmartUnmarshal[[]resource.Resource](res)
	if err != nil {
		return m, err
	}

	for _, r := range *unpacked {
		m[r.ID] = r
	}

	resourceMap = m

	return resourceMap, nil
}
