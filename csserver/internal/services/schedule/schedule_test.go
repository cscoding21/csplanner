package schedule_test

import (
	"csserver/internal/services/project"
	"csserver/internal/services/schedule"
	"csserver/internal/testobjects"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestScheduleProject(t *testing.T) {
	EXPECTED_WEEKS := 5
	EXPECTED_EXCEPTIONS := 1

	testCases := []struct {
		taskID string
	}{
		/* high skilled resource with no comms cost and 120 hour base */
		{taskID: "task:1"},
		{taskID: "task:2"},
		{taskID: "task:3"},
		{taskID: "task:4"},
		{taskID: "task:5"},
		// {taskID: "task:6"}, //---this task can't be scheduled due to no resource assignment
	}

	proj := testobjects.GetTestProjectWithCalcsDone("project:1")
	rm := testobjects.GetResourceMap()
	ram := schedule.NewResourceAllocationMapFromResourceMap(rm)

	for _, m := range proj.ProjectMilestones {
		fmt.Printf("Milestone: %s\n", m.Phase.Name)

		for _, t := range m.Tasks {
			fmt.Printf(" * %s | %v | %v | %v\n", t.Name, t.Calculated.ActualizedHoursToComplete, t.ResourceIDs, t.RequiredSkillID)
		}
	}

	result, err := schedule.ScheduleProjectAlgo(&proj, time.Now(), ram)
	if err != nil {
		t.Error(err)
	}

	if result.ProjectActivityWeeks == nil {
		t.Error("expected non-nil result")
		return
	}

	fmt.Println("***************************************")

	for _, r := range result.ProjectActivityWeeks {
		fmt.Printf("Week Ending: %v\n", r.End)

		for _, a := range r.Activities {
			fmt.Printf("  -- Task: %s | Resource: %s | Hours: %v\n", a.TaskName, a.ResourceName, a.HoursSpent)
		}
	}

	fmt.Println("---------- Exceptions -------------")
	for _, e := range result.Exceptions {
		fmt.Printf(" - %v\n", e)
	}

	if len(result.ProjectActivityWeeks) != EXPECTED_WEEKS {
		t.Errorf("Unexpected number of weeks calculated.  Expected %v - got %v", EXPECTED_WEEKS, len(result.ProjectActivityWeeks))
	}

	if len(result.Exceptions) != EXPECTED_EXCEPTIONS {
		t.Errorf("Unexpected number of exceptions.  Expected %v - got %v", EXPECTED_EXCEPTIONS, len(result.Exceptions))
	}

	for _, tc := range testCases {
		taskProj := getTaskHoursFromProject(proj, tc.taskID)
		taskSched := getTaskHoursFromSchedule(result, tc.taskID)

		if taskProj != taskSched {
			t.Errorf("Project task hours (%v) not equal to scheduled task hours (%v) for task '%v'", taskProj, taskSched, tc.taskID)
		} else {
			fmt.Printf("- Task %s hours (%v) equal scheduled hours (%v)\n", tc.taskID, taskProj, taskSched)
		}
	}
}

func TestGetScheduleItems(t *testing.T) {
	proj := testobjects.GetTestProjectWithCalcsDone("project:1")
	expectedLength := 2

	items := schedule.GetScheduleItems(&proj)

	if len(items) != expectedLength {
		t.Errorf("no items extracted.  expected %v, got %v", expectedLength, len(items))
	}
}

func getTaskHoursFromProject(proj project.Project, taskID string) int {
	hours := 0

	for _, m := range proj.ProjectMilestones {
		for _, t := range m.Tasks {
			if strings.EqualFold(*t.ID, taskID) {
				hours += t.Calculated.ActualizedHoursToComplete
			}
		}
	}

	return hours
}

func getTaskHoursFromSchedule(schedule schedule.Schedule, taskID string) int {
	hours := 0

	for _, w := range schedule.ProjectActivityWeeks {
		for _, a := range w.Activities {
			if strings.EqualFold(a.TaskID, taskID) {
				hours += a.HoursSpent
			}
		}
	}

	return hours
}
