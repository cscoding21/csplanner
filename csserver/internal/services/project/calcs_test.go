package project

import (
	"testing"
)

func TestCalculateStatsForTask(t *testing.T) {
	proj := GetTestProject()
	resources := GetTestResources()
	org := GetTestOrganization()

	task := proj.ProjectMilestones[0].Tasks[0]

	calculateStatsForTask(task, org, resources)

	if task.Calculated.ActualizedHoursToComplete != 150 {
		t.Error("expected value not correct")
	}
}
