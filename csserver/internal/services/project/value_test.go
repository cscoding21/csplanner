package project_test

import (
	"csserver/internal/services/project"
	"csserver/internal/testobjects"
	"testing"

	"github.com/cscoding21/csmap/utils"
)

func TestDeleteValueLineFromProject(t *testing.T) {
	proj := testobjects.GetTestProject("project:1")

	currentVLCount := len(proj.ProjectValue.ProjectValueLines)
	expected := currentVLCount - 1

	updatedProject := project.DeleteValueLineProjectGraph(proj, "projectvalue:1")

	updatedVLCount := len(updatedProject.ProjectValue.ProjectValueLines)

	if updatedVLCount != expected {
		t.Errorf("DeleteValueLineProjectGraph(%v) returned %v, expected %v", proj.ID, updatedVLCount, expected)
	}
}

func TestUpdateProjectValueLine_Update(t *testing.T) {
	proj := testobjects.GetTestProject("project:1")

	currentVLCount := len(proj.ProjectValue.ProjectValueLines)

	//---no change in number of value lines
	expected := currentVLCount

	updatedProject := project.UpdateProjectValueLineGraph(proj,
		project.ProjectValueLine{
			ID:             utils.ValToRef("projectvalue:1"),
			FundingSource:  "customer",
			ValueCategory:  "revenue",
			YearOneValue:   100.0,
			YearTwoValue:   100.0,
			YearThreeValue: 100.0,
			YearFourValue:  100.0,
			YearFiveValue:  100.0,
		})

	updatedVLCount := len(updatedProject.ProjectValue.ProjectValueLines)

	if updatedVLCount != expected {
		t.Errorf("UpdateProjectValueLineGraph returned %v, expected %v", updatedVLCount, expected)
	}
}

func TestUpdateProjectValueLine_New(t *testing.T) {
	proj := testobjects.GetTestProject("project:1")

	currentVLCount := len(proj.ProjectValue.ProjectValueLines)

	//---no change in number of value lines
	expected := currentVLCount + 1

	updatedProject := project.UpdateProjectValueLineGraph(proj,
		project.ProjectValueLine{
			ID:             utils.ValToRef("projectvalue:9999"),
			FundingSource:  "customer",
			ValueCategory:  "revenue",
			YearOneValue:   100.0,
			YearTwoValue:   100.0,
			YearThreeValue: 100.0,
			YearFourValue:  100.0,
			YearFiveValue:  100.0,
		})

	updatedVLCount := len(updatedProject.ProjectValue.ProjectValueLines)

	if updatedVLCount != expected {
		t.Errorf("UpdateProjectValueLineGraph returned %v, expected %v", updatedVLCount, expected)
	}

	newValueLine := updatedProject.ProjectValue.ProjectValueLines[updatedVLCount-1]

	if *newValueLine.ID != "projectvalue:9999" {
		t.Errorf("UpdateProjectValueLineGraph returned %v, expected %v", *newValueLine.ID, "projectvalue:9999")
	}
}
