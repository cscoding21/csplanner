package portfolio

import (
	"csserver/internal/calendar"
	"csserver/internal/services/resource"
	"fmt"
	"strings"
)

type ResourceAllocationMap struct {
	Portfolio      Portfolio
	WeekActivities []ResourceProjectHourAllocation
}

type ResourceProjectHourAllocation struct {
	Week                     calendar.CSWeek
	ProjectID                string
	ProjectName              string
	ResourceID               string
	ResourceName             string
	TotalResourceHours       int
	HoursAvailableForProject int
	Contention               int
}

func (ram *ResourceAllocationMap) GetResource(week calendar.CSWeek, projectID string, resourceID string) *ResourceProjectHourAllocation {
	for _, act := range ram.WeekActivities {
		if act.Week.Equal(week) && strings.EqualFold(act.ResourceID, resourceID) && strings.EqualFold(act.ProjectID, projectID) {
			return &act
		}
	}

	return nil
}

func (ram *ResourceAllocationMap) GetDistinctWeeksInFlatmap() []calendar.CSWeek {
	m := make(map[calendar.CSWeek]calendar.CSWeek)

	for _, act := range ram.WeekActivities {
		week := calendar.GetWeek(act.Week.Begin)

		m[week] = week
	}

	out := []calendar.CSWeek{}

	for k := range m {
		out = append(out, k)
	}

	return out
}

// GetResourceAllocationMap return a resource allocation map
func GetResourceAllocationMap(
	rm map[string]resource.Resource,
	portfolio Portfolio) (ResourceAllocationMap, error) {

	ram := ResourceAllocationMap{
		Portfolio: portfolio,
	}

	flattened, err := flattenPortfolio(portfolio, rm)
	if err != nil {
		return ram, err
	}

	ram.WeekActivities = flattened

	return ram, nil
}

func flattenPortfolio(p Portfolio, rm map[string]resource.Resource) ([]ResourceProjectHourAllocation, error) {
	out := []ResourceProjectHourAllocation{}

	for _, proj := range p.Schedule {
		for _, week := range proj.ProjectActivityWeeks {
			for _, act := range week.Activities {
				resource, ok := rm[act.ResourceID]
				if !ok {
					return out, fmt.Errorf("resource %s not found in pool", act.ResourceID)
				}

				ram := ResourceProjectHourAllocation{
					ProjectID:          proj.ProjectID,
					ProjectName:        proj.ProjectName,
					ResourceID:         act.ResourceID,
					ResourceName:       act.ResourceName,
					TotalResourceHours: resource.AvailableHoursPerWeek,
					Week: calendar.CSWeek{
						WeekNumber: week.WeekNumber,
						Begin:      week.Begin,
						End:        week.End,
						Year:       week.Year,
					},
				}

				out = append(out, ram)
			}
		}
	}

	calculateProjectWeekHours(&out)

	return out, nil
}

func calculateProjectWeekHours(flatmap *[]ResourceProjectHourAllocation) {
	for i, mapItem := range *flatmap {
		pm := make(map[string]string)
		week := mapItem.Week
		rid := mapItem.ResourceID
		contention := 0

		acts := findResourceUtilizationForWeek(week, rid, flatmap)
		if len(acts) == 0 {
			continue
		}

		for _, a := range acts {
			_, has := pm[a.ProjectID]
			if !has {
				pm[a.ProjectID] = a.ProjectID
				contention++
			}
		}

		hoursDivided := mapItem.TotalResourceHours / contention

		(*flatmap)[i].Contention = contention
		(*flatmap)[i].HoursAvailableForProject = hoursDivided
	}

}

func findResourceUtilizationForWeek(week calendar.CSWeek, resourceID string, flatmap *[]ResourceProjectHourAllocation) []ResourceProjectHourAllocation {
	out := []ResourceProjectHourAllocation{}

	for _, act := range *flatmap {
		if act.Week == week {
			if act.ResourceID == resourceID {
				out = append(out, act)
			}
		}
	}

	return out
}
