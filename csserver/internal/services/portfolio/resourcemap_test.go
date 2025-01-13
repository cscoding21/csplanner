package portfolio

import (
	"fmt"
	"testing"
	"time"
)

func TestGetResourceAllocationMap(t *testing.T) {
	rm := getResourceMap()
	portfolio := Portfolio{Schedule: []ProjectSchedule{}}

	testProject := getTestProjectWithCalcsDone("project:1")
	schedule, err := scheduleProject(&testProject, time.Date(2025, time.January, 21, 0, 0, 0, 0, time.UTC), rm)
	if err != nil {
		t.Error(err)
	}

	portfolio.Schedule = append(portfolio.Schedule, schedule)

	allocationMap, err := GetResourceAllocationMap(rm, portfolio)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(allocationMap)
}

func TestFlattenPortfolio(t *testing.T) {
	rm := getResourceMap()
	portfolio := getTestPortfolio2()

	allocationMap, err := flattenPortfolio(portfolio, rm)
	if err != nil {
		t.Error(err)
	}

	for _, m := range allocationMap {
		fmt.Println(m)
	}
}

func getTestPortfolio2() Portfolio {
	rm := getResourceMap()
	portfolio := Portfolio{Schedule: []ProjectSchedule{}}

	testProject := getTestProjectWithCalcsDone("project:1")
	startDate := time.Date(2025, time.January, 21, 0, 0, 0, 0, time.UTC)
	schedule, err := scheduleProject(&testProject, startDate, rm)
	if err != nil {
		panic(err)
	}

	portfolio.Schedule = append(portfolio.Schedule, schedule)

	testProject2 := getTestProjectWithCalcsDone("project:2")
	startDate = time.Date(2025, time.March, 21, 0, 0, 0, 0, time.UTC)
	schedule, err = scheduleProject(&testProject2, startDate, rm)
	if err != nil {
		panic(err)
	}

	portfolio.Schedule = append(portfolio.Schedule, schedule)

	return portfolio
}
