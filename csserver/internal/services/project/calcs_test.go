package project

import (
	"csserver/internal/services/resource"
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
		/* high skilled resource with no comms cost and 120 hour base */
		{milestoneIndex: 0, taskIndex: 0, expectedSkillsHours: -30, expectedCommsHours: 0, expectedActualHours: 90, expectedExceptions: 0},
		/* mid skilled resource with no comms cost and 40 hour base */
		{milestoneIndex: 0, taskIndex: 1, expectedSkillsHours: 0, expectedCommsHours: 0, expectedActualHours: 40, expectedExceptions: 0},
		/* low skilled resource with no comms cost and 40 hour base */
		{milestoneIndex: 0, taskIndex: 2, expectedSkillsHours: 10, expectedCommsHours: 0, expectedActualHours: 50, expectedExceptions: 0},
		/* high & mid skilled resource with comms cost for 2 resources and 120 hour base */
		{milestoneIndex: 1, taskIndex: 0, expectedSkillsHours: -30, expectedCommsHours: 12, expectedActualHours: 102, expectedExceptions: 0},
		/* resource exception: allocated resource does not have required skill */
		{milestoneIndex: 1, taskIndex: 1, expectedSkillsHours: 0, expectedCommsHours: 0, expectedActualHours: 40, expectedExceptions: 1},
	}

	proj := GetTestProject()
	resources := GetTestResources()
	org := GetTestOrganization()
	rm := ConvertResourceSliceToMap(resources)

	for _, tc := range testCases {
		task := proj.ProjectMilestones[tc.milestoneIndex].Tasks[tc.taskIndex]

		calculateStatsForTask(task, org, rm)

		if task.Calculated.SkillsHourAdjustment != tc.expectedSkillsHours {
			t.Errorf("expected skills hours value not correct.  want %v - got %v", tc.expectedSkillsHours, task.Calculated.SkillsHourAdjustment)
		}

		if task.Calculated.CommsHourAdjustment != tc.expectedCommsHours {
			t.Errorf("expected comms hours value not correct.  want %v - got %v", tc.expectedCommsHours, task.Calculated.CommsHourAdjustment)
		}

		if task.Calculated.ActualizedHoursToComplete != tc.expectedActualHours {
			t.Errorf("expected value not correct.  want %v - got %v", tc.expectedActualHours, task.Calculated.ActualizedHoursToComplete)
		}

		if len(task.Calculated.Exceptions) != tc.expectedExceptions {
			t.Errorf("resourse expected exception to occur but was not found")
		}

		if task.Calculated.ActualizedHoursToComplete != (task.HourEstimate + task.Calculated.SkillsHourAdjustment + task.Calculated.CommsHourAdjustment) {
			t.Errorf("actualized hours do not equal sum of estimate + skills adjustment + comms adjustment")
		}
	}
}

func ConvertResourceSliceToMap(resources []resource.Resource) map[string]resource.Resource {
	m := make(map[string]resource.Resource)
	for _, r := range resources {
		m[r.ID] = r
	}

	return m
}
