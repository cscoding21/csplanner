package portfolio

import (
	"csserver/internal/calendar"
	"csserver/internal/services/resource"
	"fmt"
)

type ResourceAllocationMap struct {
	//week -> project -> resource
	Map map[calendar.CSWeek]map[string]map[string]ResourceProjectHourAllocation
}

type ResourceProjectHourAllocation struct {
	ResourceID   string
	ResourceName string
	Hours        int
}

func (ram *ResourceAllocationMap) GetResource(week calendar.CSWeek, projectID string, resourceID string) ResourceProjectHourAllocation {
	byWeek := ram.Map[week]

	byProject := byWeek[projectID]
	r := byProject[resourceID]

	return r
}

// GetResourceAllocationMap return a resource allocation map
func GetResourceAllocationMap(
	rm map[string]resource.Resource,
	portfolio Portfolio) (ResourceAllocationMap, error) {

	ram := ResourceAllocationMap{
		Map: make(map[calendar.CSWeek]map[string]map[string]ResourceProjectHourAllocation),
	}

	rangeStart, rangeEnd := portfolio.GetDateRange()
	weeks := calendar.GetWeeks(rangeStart, rangeEnd)

	for _, w := range weeks {
		ram.Map[w] = make(map[string]map[string]ResourceProjectHourAllocation)

		for _, s := range portfolio.Schedule {
			fmt.Println(s)
		}
	}

	return ram, nil
}
