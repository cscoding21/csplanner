package project

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/services/activity"
	"csserver/internal/services/organization"
	"csserver/internal/services/project/ptypes/projectstatus"
	"csserver/internal/services/projecttemplate"
	"csserver/internal/services/resource"
	"csserver/internal/utils"

	"github.com/cscoding21/csval/validate"
)

var (
	ActivityNewProject          = activity.ActivityDef{Type: "new-project", Template: "The project '%s' was created by %s"}
	ActivityUpdateProject       = activity.ActivityDef{Type: "update-project", Template: "The project '%s' was updated by %s"}
	ActivityProjectStatusChange = activity.ActivityDef{Type: "project-status-change", Template: ""}
)

// SaveProject saves a project in the system and updates all stored calculations
func (s *ProjectService) SaveProject(
	ctx context.Context,
	pro Project,
	resourceMap map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization) (common.UpdateResult[*common.BaseModel[Project]], error) {
	//---TODO: update to proper validation
	//val := pro.Validate()
	val := validate.NewSuccessValidationResult()

	if len(pro.ID) == 0 {
		return s.newProject(ctx, pro, resourceMap, roleMap, org, val)
	}

	lastProject, err := s.GetProjectByID(ctx, pro.ID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}
	if lastProject == nil {
		//---existing project does not exist...create
		return s.newProject(ctx, pro, resourceMap, roleMap, org, val)
	}

	lastProject.Data.ProjectBasics = pro.ProjectBasics
	lastProject.Data.ProjectCost = pro.ProjectCost
	lastProject.Data.ProjectDaci = pro.ProjectDaci

	lastProject.Data.ProjectValue.DiscountRate = pro.ProjectValue.DiscountRate

	lastProject.Data.PerformAllCalcs(org, resourceMap, roleMap)

	return s.UpdateProject(ctx, lastProject.Data)
}

func (s *ProjectService) newProject(
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

	return s.UpdateProject(ctx, pro)
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

	return s.SaveProject(ctx, *pro, rm, roleMap, org)
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

	newState := stateMachineMap[status]

	val := newState.Can(&p.Data)

	if !val.Pass && !force {
		return common.NewFailingUpdateResultWithVal[*common.BaseModel[Project]](nil, val)
	}

	p.Data.ProjectStatusBlock.Status = newState.State

	return s.UpdateProject(ctx, p.Data)
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
