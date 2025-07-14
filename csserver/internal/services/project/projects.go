package project

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/services/organization"
	"csserver/internal/services/project/ptypes/projectstatus"
	"csserver/internal/services/projecttemplate"
	"csserver/internal/services/resource"
	"csserver/internal/utils"
	"fmt"

	"github.com/cscoding21/csval/validate"

	log "github.com/sirupsen/logrus"
)

// SaveProject saves a project in the system and updates all stored calculations
func (s *ProjectService) SaveProject(
	ctx context.Context,
	pro Project,
	resourceMap map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization,
	eventName *string) (common.UpdateResult[*common.BaseModel[Project]], error) {
	//---TODO: update to proper validation
	// val := pro.Validate()
	val := validate.NewSuccessValidationResult()
	log.Debugf("project validation: %v", val)

	if len(pro.ID) == 0 {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, fmt.Errorf("project id not known.  Save failed"))
	}

	if eventName == nil {
		eventName = utils.ValToRef("updated")
	}

	lastProject, err := s.GetProjectByID(ctx, pro.ID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}
	if lastProject == nil {
		//---existing project does not exist...create
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, fmt.Errorf("project with id %s not found.  Save failed", pro.ID))
	}

	lastProject.Data.ProjectBasics = pro.ProjectBasics
	lastProject.Data.ProjectCost = pro.ProjectCost
	lastProject.Data.ProjectDaci = pro.ProjectDaci

	lastProject.Data.ProjectValue.DiscountRate = pro.ProjectValue.DiscountRate

	lastProject.Data.PerformAllCalcs(org, resourceMap, roleMap)

	resp, err := s.UpdateProject(ctx, lastProject.Data)

	updated := *resp.Object
	s.pubsub.StreamPublish(ctx,
		string(ProjectIdentifier),
		"project",
		*eventName,
		map[string]any{
			"id":   updated.ID,
			"name": updated.Data.ProjectBasics.Name,
		})

	return resp, err
}

func (s *ProjectService) SaveTestProject(
	ctx context.Context,
	pro Project,
	resourceMap map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization,
	val validate.ValidationResult) (common.UpdateResult[*common.BaseModel[Project]], error) {

	pro.ProjectStatusBlock.Status = projectstatus.NewProject

	updateResult, err := s.CreateProject(ctx, pro)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}

	updateObj := *updateResult.Object
	pro = updateObj.Data

	pro.PerformAllCalcs(org, resourceMap, roleMap)

	resp, err := s.UpdateProject(ctx, pro)

	return resp, err
}

// CreateNewProject adds a new project to the portfolio
func (s *ProjectService) CreateNewProject(
	ctx context.Context,
	basics ProjectBasics,
	template projecttemplate.Projecttemplate,
	rm map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization) (common.UpdateResult[*common.BaseModel[Project]], error) {

	newID := utils.GenerateBase64UUID()

	newProj := Project{
		ControlFields: common.ControlFields{
			ID: newID,
		},
		ProjectBasics: &basics,
		ProjectStatusBlock: &ProjectStatusBlock{
			Status: projectstatus.NewProject,
		},
		ProjectValue: &ProjectValue{
			DiscountRate: org.Defaults.DiscountRate,
			Calculated:   ProjectValueCalculatedData{},
		},
		ProjectCost: &ProjectCost{
			Calculated: ProjectCostCalculatedData{},
		},
		ProjectDaci: &ProjectDaci{
			DriverIDs:      []*string{},
			ApproverIDs:    []*string{},
			ContributorIDs: []*string{},
			InformedIDs:    []*string{},
		},
		ProjectFeatures: []*ProjectFeature{},
		Calculated:      ProjectCalculatedData{},
	}

	updateResult, err := s.CreateProject(ctx, newProj)
	if err != nil {
		return updateResult, err
	}

	pro := common.UpwrapFromUpdateResult(updateResult)

	updateResultWithMS, err := s.SetProjectMilestonesFromTemplate(ctx, pro.ID, template, rm, roleMap, org)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}

	pro = common.UpwrapFromUpdateResult(updateResultWithMS)

	return s.SaveProject(ctx, *pro, rm, roleMap, org, utils.ValToRef("created"))
}

// SetProjectStatus update the status of a project by adhering to the rules of the state machine
func (s *ProjectService) SetProjectStatus(
	ctx context.Context,
	projectID string,
	status projectstatus.ProjectState,
	force bool) (common.UpdateResult[*common.BaseModel[Project]], error) {

	p, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}

	oldState := stateMachineMap[p.Data.ProjectStatusBlock.Status]
	newState := stateMachineMap[status]

	val := newState.Can(&p.Data)

	if !val.Pass && !force {
		return common.NewFailingUpdateResultWithVal[*common.BaseModel[Project]](nil, val)
	}

	p.Data.ProjectStatusBlock.Status = newState.State

	resp, err := s.UpdateProject(ctx, p.Data)
	up := *resp.Object
	s.pubsub.StreamPublish(ctx,
		string(ProjectIdentifier),
		"state",
		"updated",
		map[string]any{
			"id":         up.ID,
			"name":       up.Data.ProjectBasics.Name,
			"start_date": up.Data.ProjectBasics.StartDate,
			"from_state": oldState.State,
			"to_state":   newState.State,
		})

	return resp, err
}

// CheckProjectStatusChange performa dry run of a status change to see if it would go through successfully
func (s *ProjectService) CheckProjectStatusChange(ctx context.Context, projectID string, status projectstatus.ProjectState) (validate.ValidationResult, error) {
	p, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		ev := validate.NewFailingValidationResult(validate.NewValidationMessage("", err.Error()))
		return ev, err
	}

	newState := stateMachineMap[status]

	return newState.Can(&p.Data), nil
}

// GetStatusTransitionDetails populate information about the project status state machine
func (s *ProjectService) GetStatusTransitionDetails(p *Project) {
	stateInfo := stateMachineMap[p.ProjectStatusBlock.Status]

	for _, si := range stateInfo.NextValidStates {
		pst := ProjectStatusTransition{
			NextState: si,
		}

		potentialStateInfo := stateMachineMap[si]
		result := potentialStateInfo.Can(p)
		pst.CheckResults = result
		pst.CanEnter = result.Pass

		p.ProjectStatusBlock.AllowedNextStates = append(p.ProjectStatusBlock.AllowedNextStates, &pst)
	}
}
