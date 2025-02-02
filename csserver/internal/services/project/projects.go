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

	log "github.com/sirupsen/logrus"
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

	lastProject.ProjectBasics = pro.ProjectBasics
	lastProject.ProjectCost = pro.ProjectCost
	lastProject.ProjectDaci = pro.ProjectDaci

	lastProject.ProjectValue.DiscountRate = pro.ProjectValue.DiscountRate
	lastProject.ProjectValue.IsCapitalized = pro.ProjectValue.IsCapitalized

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
	pro.ProjectStatusBlock.Status = projectstatus.NewProject
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
func (s *ProjectService) SetProjectStatus(
	ctx context.Context,
	projectID string,
	status projectstatus.ProjectState,
	force bool) (common.UpdateResult[Project], error) {

	p, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		ev := validate.NewFailingValidationResult(validate.NewValidationMessage("", err.Error()))
		return common.NewUpdateResult(&ev, p), err
	}

	newState := stateMachineMap[status]

	val := newState.Can(p)

	if !val.Pass && !force {
		return common.NewUpdateResult(&val, p), fmt.Errorf("illegal state transition")
	}

	p.ProjectStatusBlock.Status = newState.State

	return s.UpdateProject(ctx, p)
}

// CheckProjectStatusChange performa dry run of a status change to see if it would go through successfully
func (s *ProjectService) CheckProjectStatusChange(ctx context.Context, projectID string, status projectstatus.ProjectState) (validate.ValidationResult, error) {
	p, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		ev := validate.NewFailingValidationResult(validate.NewValidationMessage("", err.Error()))
		return ev, err
	}

	newState := stateMachineMap[status]

	return newState.Can(p), nil
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

// GetStatusTransitionDetails populate information about the project status state machine
func (s *ProjectService) GetStatusTransitionDetails(p *Project) {
	stateInfo := stateMachineMap[p.ProjectStatusBlock.Status]

	log.Warn(stateInfo)
	for _, si := range stateInfo.NextValidStates {
		pst := ProjectStatusTransition{
			NextState: si,
		}

		potentialStateInfo := stateMachineMap[si]
		result := potentialStateInfo.Can(p)
		pst.CheckResults = result
		pst.CanEnter = result.Pass

		log.Warn(pst)

		p.ProjectStatusBlock.AllowedNextStates = append(p.ProjectStatusBlock.AllowedNextStates, &pst)
	}
}
