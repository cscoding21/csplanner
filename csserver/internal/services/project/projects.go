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
	// lastProject, err := s.GetProjectByID(ctx, pro.ID)
	// if err != nil {
	// 	return common.NewUpdateResult(&val, &pro), err
	// }

	if len(pro.ID) == 0 {
		proj, err := s.CreateProject(ctx, &pro)
		if err != nil {
			return common.NewUpdateResult(&val, &pro), err
		}

		pro = *proj.Object
	}

	pro.GetProjectNPV()
	pro.GetProjectIRR()
	pro.GetProjectInitialCost()

	return s.UpdateProject(ctx, &pro)
}
