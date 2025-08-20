package project

import (
	"context"
	"slices"

	"csserver/internal/common"
	"csserver/internal/services/organization"
	"csserver/internal/services/resource"

	"csserver/internal/utils"

	"github.com/google/uuid"
)

// DeleteTaskFromProject if the feature exists in the project object graph...remove it
func (s *ProjectService) DeleteFeatureFromProject(
	ctx context.Context,
	projectID string,
	featureID string,
	resourceMap map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization,
) (common.UpdateResult[*common.BaseModel[Project]], error) {

	result, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}

	project := result.Data

	updatedProject, deletedFeature := DeleteFeatureFromProjectGraph(project, featureID)
	updatedProject.PerformAllCalcs(org, resourceMap, roleMap)

	s.pubsub.StreamPublish(ctx,
		string(ProjectIdentifier),
		"feature",
		"deleted",
		map[string]any{
			"id":           featureID,
			"project_id":   projectID,
			"project_name": project.ProjectBasics.Name,
			"name":         deletedFeature.Name,
		})

	return s.UpdateProject(ctx, updatedProject)
}

// DeleteTaskFromProjectGraph if the feature exists in the project, remove it.
func DeleteFeatureFromProjectGraph(project Project, featureID string) (Project, ProjectFeature) {
	var deletedFeature ProjectFeature
	for i, f := range project.ProjectFeatures {
		if f.ID != nil && *f.ID == featureID {
			deletedFeature = *project.ProjectFeatures[i]
			project.ProjectFeatures = slices.Delete(project.ProjectFeatures, i, i+1)

			break
		}
	}

	return project, deletedFeature
}

// UpdateProjectFeature if the feature exists in the project object graph...update it.  Otherwise, add it
func (s *ProjectService) UpdateProjectFeature(
	ctx context.Context,
	projectID string,
	feature ProjectFeature,
	resourceMap map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization,
) (common.UpdateResult[*common.BaseModel[Project]], error) {
	projectResult, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}

	dg := ""
	op := "created"

	project := projectResult.Data
	updatedProject, existingFeature := UpdateProjectFeatureGraph(project, feature)
	updatedProject.PerformAllCalcs(org, resourceMap, roleMap)

	if existingFeature != nil {
		dg = utils.GetDiffGraph(*existingFeature, feature)
		op = "updated"
	}

	s.pubsub.StreamPublish(ctx,
		string(ProjectIdentifier),
		"feature",
		op,
		map[string]any{
			"id":           feature.ID,
			"project_id":   projectID,
			"project_name": project.ProjectBasics.Name,
			"name":         feature.Name,
			"diffs":        dg,
		})

	return s.UpdateProject(ctx, updatedProject)
}

// updateProjectFeatureGraph if the feature exists in the project object graph...update it.  Otherwise, add it
func UpdateProjectFeatureGraph(project Project, feature ProjectFeature) (Project, *ProjectFeature) {
	for i, f := range project.ProjectFeatures {
		if f.ID != nil && *f.ID == *feature.ID {
			project.ProjectFeatures[i] = &feature
			return project, f
		}
	}

	//---tis a new feature
	if len(*feature.ID) == 0 {
		feature.ID = utils.ValToRef(uuid.New().String())
	}

	project.ProjectFeatures = append(project.ProjectFeatures, &feature)

	return project, nil
}
