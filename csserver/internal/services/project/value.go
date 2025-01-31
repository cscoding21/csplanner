package project

import (
	"context"
	"csserver/internal/common"
	"slices"

	"github.com/cscoding21/csmap/utils"
	"github.com/google/uuid"
)

// DeleteValueLineProject if the value like exists in the project object graph...remove it
func (s *ProjectService) DeleteValueLineProject(ctx context.Context, projectID string, valueLineID string) (*common.UpdateResult[Project], error) {
	project, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.HandleReturnWithValue[common.UpdateResult[Project]](nil, err)
	}

	updatedProject := DeleteValueLineProjectGraph(*project, valueLineID)
	pro, err := s.UpdateProject(ctx, &updatedProject)
	return &pro, err
}

// DeleteValueLineProjectGraph if the value line exists in the project, remove it.
func DeleteValueLineProjectGraph(project Project, valueLineID string) Project {
	for i, f := range project.ProjectValue.ProjectValueLines {
		if f.ID != nil && *f.ID == valueLineID {
			project.ProjectValue.ProjectValueLines = slices.Delete(project.ProjectValue.ProjectValueLines, i, i+1)

			break
		}
	}

	project.AggregateProjectValueLines()
	project.GetProjectNPV()
	project.GetProjectIRR()

	return project
}

// UpdateProjectValueLine if the value line exists in the project object graph...update it.  Otherwise, add it
func (s *ProjectService) UpdateProjectValueLine(ctx context.Context, projectID string, valueLine ProjectValueLine) (*common.UpdateResult[Project], error) {
	project, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.HandleReturnWithValue[common.UpdateResult[Project]](nil, err)
	}

	updatedProject := UpdateProjectValueLineGraph(*project, valueLine)
	updatedProject.AggregateProjectValueLines()
	updatedProject.GetProjectNPV()
	updatedProject.GetProjectIRR()

	pro, err := s.UpdateProject(ctx, &updatedProject)
	return common.HandleReturnWithValue[common.UpdateResult[Project]](&pro, err)
}

// updateProjectValueLineGraph if the value line exists in the project object graph...update it.  Otherwise, add it
func UpdateProjectValueLineGraph(project Project, valueLine ProjectValueLine) Project {
	for i, f := range project.ProjectValue.ProjectValueLines {
		if f.ID != nil && *f.ID == *valueLine.ID {
			project.ProjectValue.ProjectValueLines[i] = &valueLine
			return project
		}
	}

	//---tis a new value line
	if len(*valueLine.ID) == 0 {
		valueLine.ID = utils.ValToRef(uuid.New().String())
	}

	project.ProjectValue.ProjectValueLines = append(project.ProjectValue.ProjectValueLines, &valueLine)

	return project
}
