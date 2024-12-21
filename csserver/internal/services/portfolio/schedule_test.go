package portfolio

import (
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"fmt"
	"testing"
	"time"
)

func TestScheduleProject(t *testing.T) {
	proj := getTestProjectWithCalcsDone()
	rm := getResourceMap()

	for _, m := range proj.ProjectMilestones {
		fmt.Printf("Milestone: %s\n", m.Phase.Name)

		for _, t := range m.Tasks {
			fmt.Printf(" * %s | %v\n", t.Name, t.Calculated.ActualizedHoursToComplete)
		}
	}

	result, err := ScheduleProject(&proj, time.Now(), rm)
	if err != nil {
		t.Error(err)
	}

	if result.ProjectActivityWeeks == nil {
		t.Error("expected non-nil result")
		return
	}

	fmt.Println("***************************************")

	for _, r := range *result.ProjectActivityWeeks {
		fmt.Printf("Week Ending: %v\n", r.Week)

		for _, a := range r.Activities {
			fmt.Printf("  -- Task: %s | Resource: %s | Hours: %v\n", a.TaskName, a.ResourceName, a.HoursSpent)
		}
	}

	fmt.Println("---------- Exceptions -------------")
	for _, e := range result.Exceptions {
		fmt.Printf(" - %s\n", e)
	}

	fmt.Println(proj.ProjectBasics.Name)
}

func TestGetScheduleItems(t *testing.T) {
	proj := getTestProjectWithCalcsDone()
	expectedLength := 2

	items := getScheduleItems(&proj)

	if len(items) != expectedLength {
		t.Errorf("no items extracted.  expected %v, got %v", expectedLength, len(items))
	}
}

func TestGetResourceHoursForWeek(t *testing.T) {
	rm := getResourceMap()
	expectedLength := 5

	hoursMap := getResourceHoursForWeek(rm)

	if len(hoursMap) != expectedLength {
		t.Errorf("no items extracted.  expected %v, got %v", expectedLength, len(hoursMap))
	}
}

func getTestProjectWithCalcsDone() project.Project {
	proj := GetTestProject()
	rm := getResourceMap()
	org := GetTestOrganization()

	proj.CalculateProjectTasksStats(org, rm)

	return proj
}

func getResourceMap() map[string]resource.Resource {
	resources := GetTestResources()
	return ConvertResourceSliceToMap(resources)
}
