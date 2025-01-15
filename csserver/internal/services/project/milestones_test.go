package project_test

import (
	"testing"

	"csserver/internal/services/project"
	"csserver/internal/testobjects"
	"csserver/internal/utils"

	log "github.com/sirupsen/logrus"
)

func TestUpdateProjectTask(t *testing.T) {
	proj := testobjects.GetTestProject("project:1")

	for i := range proj.ProjectMilestones {
		log.Infof("%v - %s", *proj.ProjectMilestones[i].ID, proj.ProjectMilestones[i].Phase.Name)
		for j := range proj.ProjectMilestones[i].Tasks {
			log.Infof("--- %s - %v", *proj.ProjectMilestones[i].Tasks[j].ID, *proj.ProjectMilestones[i].Tasks[j])
		}
	}

	updatedProject := project.UpdateTaskFromProjectGraph(proj, "milestone:1", project.ProjectMilestoneTask{
		ID:           utils.ValToRef("task:1"),
		HourEstimate: 30,
		Name:         "Updated Project Task",
		Status:       "accepted",
	})

	for i := range updatedProject.ProjectMilestones {
		log.Infof("%v - %s", *updatedProject.ProjectMilestones[i].ID, updatedProject.ProjectMilestones[i].Phase.Name)
		for j := range updatedProject.ProjectMilestones[i].Tasks {
			log.Infof("--- %s - %v", *updatedProject.ProjectMilestones[i].Tasks[j].ID, *updatedProject.ProjectMilestones[i].Tasks[j])
		}
	}

	if len(proj.ProjectMilestones[0].Tasks) != len(updatedProject.ProjectMilestones[0].Tasks) {
		t.Errorf("updateTaskFromProjectGraph returned %v, expected %v", len(updatedProject.ProjectMilestones[0].Tasks), len(proj.ProjectMilestones[0].Tasks))
	}
}

func TestAddProjectTask(t *testing.T) {
	proj := testobjects.GetTestProject("project:1")

	expectedCount := len(proj.ProjectMilestones[0].Tasks) + 1

	for i := range proj.ProjectMilestones {
		log.Infof("%v - %s", *proj.ProjectMilestones[i].ID, proj.ProjectMilestones[i].Phase.Name)
		for j := range proj.ProjectMilestones[i].Tasks {
			log.Infof("--- %s - %v", *proj.ProjectMilestones[i].Tasks[j].ID, *proj.ProjectMilestones[i].Tasks[j])
		}
	}

	updatedProject := project.UpdateTaskFromProjectGraph(proj, "milestone:1", project.ProjectMilestoneTask{
		HourEstimate: 30,
		Name:         "Added Project Task",
		Status:       "accepted",
	})

	for i := range updatedProject.ProjectMilestones {
		log.Infof("%v - %s", *updatedProject.ProjectMilestones[i].ID, updatedProject.ProjectMilestones[i].Phase.Name)
		for j := range updatedProject.ProjectMilestones[i].Tasks {
			log.Infof("--- %s - %v", *updatedProject.ProjectMilestones[i].Tasks[j].ID, *updatedProject.ProjectMilestones[i].Tasks[j])
		}
	}

	if expectedCount != len(updatedProject.ProjectMilestones[0].Tasks) {
		t.Errorf("updateTaskFromProjectGraph returned %v, expected %v", len(updatedProject.ProjectMilestones[0].Tasks), expectedCount)
	}
}
