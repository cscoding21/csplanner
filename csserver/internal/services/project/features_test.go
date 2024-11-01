package project

import (
	"testing"

	"csserver/internal/services/project/ptypes/featurepriority"
	"csserver/internal/utils"
)

func TestDeleteFeatureFromProject(t *testing.T) {
	proj, err := utils.DeepCopy(testProject)
	if err != nil {
		t.Error(err)
	}

	currentFeatureCount := len(proj.ProjectFeatures)
	expected := currentFeatureCount - 1

	updatedProject := deleteFeatureFromPrjectGraph(*proj, "projectfeature:1")

	updatedFeatureCount := len(updatedProject.ProjectFeatures)

	if updatedFeatureCount != expected {
		t.Errorf("DeleteFeatureFromProject(%v) returned %v, expected %v", testProject.ID, updatedFeatureCount, expected)
	}
}

func TestUpdateProjectFeature_Update(t *testing.T) {
	proj, err := utils.DeepCopy(testProject)
	if err != nil {
		t.Error(err)
	}

	currentFeatureCount := len(proj.ProjectFeatures)

	//---no change in number of features
	expected := currentFeatureCount

	updatedProject := updateProjectFeatureGraph(*proj,
		ProjectFeature{
			ID:          utils.ValToRef("projectfeature:1"),
			Priority:    featurepriority.Medium,
			Name:        "Topic Sentence UPDATED",
			Description: "It lays out the path UPDATED",
			Status:      "proposed"})

	updatedFeatureCount := len(updatedProject.ProjectFeatures)

	if updatedFeatureCount != expected {
		t.Errorf("updateProjectFeatureGraph returned %v, expected %v", updatedFeatureCount, expected)
	}
}

func TestUpdateProjectFeature_New(t *testing.T) {
	proj, err := utils.DeepCopy(testProject)
	if err != nil {
		t.Error(err)
	}

	currentFeatureCount := len(proj.ProjectFeatures)

	//---no change in number of features
	expected := currentFeatureCount + 1

	updatedProject := updateProjectFeatureGraph(*proj,
		ProjectFeature{
			ID:          utils.ValToRef("projectfeature:9999"),
			Priority:    featurepriority.Medium,
			Name:        "Topic Sentence ADDED",
			Description: "It lays out the path ADDED",
			Status:      "proposed"})

	updatedFeatureCount := len(updatedProject.ProjectFeatures)

	if updatedFeatureCount != expected {
		t.Errorf("updateProjectFeatureGraph returned %v, expected %v", updatedFeatureCount, expected)
	}

	newFeature := updatedProject.ProjectFeatures[updatedFeatureCount-1]

	if *newFeature.ID != "projectfeature:9999" {
		t.Errorf("updateProjectFeatureGraph returned %v, expected %v", *newFeature.ID, "projectfeature:9999")
	}
}
