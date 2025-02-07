package project_test

import (
	"csserver/internal/services/project"
	"csserver/internal/testobjects"
	"testing"
)

func TestCalculateStatsForTask(t *testing.T) {

	testCases := []struct {
		milestoneIndex      int
		taskIndex           int
		expectedSkillsHours int
		expectedCommsHours  int
		expectedActualHours int
		expectedExceptions  int
	}{
		{milestoneIndex: 0, taskIndex: 0, expectedSkillsHours: -30, expectedCommsHours: 0, expectedActualHours: 90, expectedExceptions: 0},
		{milestoneIndex: 0, taskIndex: 1, expectedSkillsHours: 0, expectedCommsHours: 0, expectedActualHours: 40, expectedExceptions: 0},
		{milestoneIndex: 0, taskIndex: 2, expectedSkillsHours: 0, expectedCommsHours: 0, expectedActualHours: 40, expectedExceptions: 0},
		{milestoneIndex: 1, taskIndex: 0, expectedSkillsHours: -30, expectedCommsHours: 12, expectedActualHours: 102, expectedExceptions: 0},
		{milestoneIndex: 1, taskIndex: 1, expectedSkillsHours: -10, expectedCommsHours: 0, expectedActualHours: 30, expectedExceptions: 0},
	}

	proj := testobjects.GetTestProject("project:1")
	resourceMap := testobjects.GetResourceMap()
	roleMap := testobjects.GetRoleMap()
	org := testobjects.GetTestOrganization()

	for i, tc := range testCases {
		task := proj.ProjectMilestones[tc.milestoneIndex].Tasks[tc.taskIndex]

		project.CalculateStatsForTask(task, org, resourceMap, roleMap)

		if task.Calculated.SkillsHourAdjustment != tc.expectedSkillsHours {
			t.Errorf("%v : expected skills hours value not correct.  want %v - got %v", i, tc.expectedSkillsHours, task.Calculated.SkillsHourAdjustment)
		}

		if task.Calculated.CommsHourAdjustment != tc.expectedCommsHours {
			t.Errorf("%v : expected comms hours value not correct.  want %v - got %v", i, tc.expectedCommsHours, task.Calculated.CommsHourAdjustment)
		}

		if task.Calculated.ActualizedHoursToComplete != tc.expectedActualHours {
			t.Errorf("%v : expected value not correct.  want %v - got %v", i, tc.expectedActualHours, task.Calculated.ActualizedHoursToComplete)
		}

		if len(task.Calculated.Exceptions) != tc.expectedExceptions {
			t.Errorf("%v : resource expected exception to occur but was not found", i)
		}

		if task.Calculated.ActualizedHoursToComplete != (task.HourEstimate + task.Calculated.SkillsHourAdjustment + task.Calculated.CommsHourAdjustment) {
			t.Errorf("%v : actualized hours do not equal sum of estimate + skills adjustment + comms adjustment", i)
		}
	}
}
