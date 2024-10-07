package project

import (
	"context"
	"csserver/internal/common"

	"github.com/cscoding21/csval/validate"
)

// SaveProject saves a project in the system and updates all stored calculations
func (s *ProjectService) SaveProject(ctx context.Context, pro Project) (common.UpdateResult[Project], error) {
	//---TODO: update to proper validation
	//val := pro.Validate()
	val := validate.NewSuccessValidationResult()

	if len(pro.ID) == 0 {
		return s.newProject(ctx, pro, val)
	}

	lastProject, err := s.GetProjectByID(ctx, pro.ID)
	if err != nil {
		//---existing project does not exist...create
		return s.newProject(ctx, pro, val)
	}

	lastProject.ProjectBasics = pro.ProjectBasics
	lastProject.ProjectValue = pro.ProjectValue
	lastProject.ProjectCost = pro.ProjectCost
	lastProject.ProjectDaci = pro.ProjectDaci

	lastProject.GetProjectInitialCost()
	lastProject.GetProjectNPV()
	lastProject.GetProjectIRR()

	return s.UpdateProject(ctx, lastProject)
}

func (s *ProjectService) newProject(ctx context.Context, pro Project, val validate.ValidationResult) (common.UpdateResult[Project], error) {
	proj, err := s.CreateProject(ctx, &pro)
	if err != nil {
		return common.NewUpdateResult(&val, &pro), err
	}

	pro = *proj.Object

	pro.GetProjectInitialCost()
	pro.GetProjectNPV()
	pro.GetProjectIRR()

	return s.UpdateProject(ctx, &pro)
}
