package project

import (
	"context"
	"slices"

	"csserver/internal/common"

	"github.com/cscoding21/csmap/utils"
	"github.com/google/uuid"
)

// DeleteTaskFromProject if the feature exists in the project object graph...remove it
func (s *ProjectService) DeleteFeatureFromProject(ctx context.Context, projectID string, featureID string) (*common.UpdateResult[Project], error) {
	project, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.HandleReturnWithValue[common.UpdateResult[Project]](nil, err)
	}

	updatedProject := DeleteFeatureFromPrjectGraph(*project, featureID)
	pro, err := s.UpdateProject(ctx, &updatedProject)
	return &pro, err
}

// DeleteTaskFromProjectGraph if the feature exists in the project, remove it.
func DeleteFeatureFromPrjectGraph(project Project, featureID string) Project {
	for i, f := range project.ProjectFeatures {
		if f.ID != nil && *f.ID == featureID {
			project.ProjectFeatures = slices.Delete(project.ProjectFeatures, i, i+1)

			break
		}
	}

	return project
}

// UpdateProjectFeature if the feature exists in the project object graph...update it.  Otherwise, add it
func (s *ProjectService) UpdateProjectFeature(ctx context.Context, projectID string, feature ProjectFeature) (*common.UpdateResult[Project], error) {
	project, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.HandleReturnWithValue[common.UpdateResult[Project]](nil, err)
	}

	updatedProject := UpdateProjectFeatureGraph(*project, feature)
	pro, err := s.UpdateProject(ctx, &updatedProject)
	return common.HandleReturnWithValue[common.UpdateResult[Project]](&pro, err)
}

// updateProjectFeatureGraph if the feature exists in the project object graph...update it.  Otherwise, add it
func UpdateProjectFeatureGraph(project Project, feature ProjectFeature) Project {
	for i, f := range project.ProjectFeatures {
		if f.ID != nil && *f.ID == *feature.ID {
			project.ProjectFeatures[i] = &feature
			return project
		}
	}

	//---tis a new feature
	if len(*feature.ID) == 0 {
		feature.ID = utils.ValToRef(uuid.New().String())
	}

	project.ProjectFeatures = append(project.ProjectFeatures, &feature)

	return project
}
