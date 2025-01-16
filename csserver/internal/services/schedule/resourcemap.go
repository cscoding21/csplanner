package schedule

import (
	"csserver/internal/calendar"
	"csserver/internal/services/resource"
	"fmt"
	"strings"
	"time"
)

type ResourceAllocationMap struct {
	// Portfolio      Portfolio
	WeekActivities []ResourceProjectHourAllocation
}

func NewResourceAllocationMapFromResourceMap(rm map[string]resource.Resource) ResourceAllocationMap {
	ram := ResourceAllocationMap{WeekActivities: []ResourceProjectHourAllocation{}}

	start := time.Now()
	end := start.Add(1 * time.Hour * 24 * 7 * 52 * 3)
	weeks := calendar.GetWeeks(start, end)

	for _, w := range weeks {
		for _, v := range rm {
			entry := ResourceProjectHourAllocation{
				Week:                     w,
				ResourceID:               v.ID,
				ResourceName:             v.Name,
				ProjectID:                "",
				ProjectName:              "Placeholder",
				Contention:               1,
				HoursAvailableForProject: v.AvailableHoursPerWeek,
				TotalResourceHours:       v.AvailableHoursPerWeek,
			}

			ram.WeekActivities = append(ram.WeekActivities, entry)
		}
	}

	return ram
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

// GetResource return resource details given a week, resourceID< and project
func (ram *ResourceAllocationMap) GetResource(week calendar.CSWeek, projectID string, resourceID string) *ResourceProjectHourAllocation {
	for _, act := range ram.WeekActivities {
		if act.Week.Equal(week) {
			if strings.EqualFold(act.ResourceID, resourceID) {
				if strings.EqualFold(act.ProjectID, projectID) || len(act.ProjectID) == 0 {
					return &act
				}
			}
		}
	}

	return nil
}

func (ram *ResourceAllocationMap) ReduceResourceProjectHours(week calendar.CSWeek, projectID string, resourceID string, hoursToSubtract int) {
	for i, act := range ram.WeekActivities {
		if act.Week.Equal(week) && strings.EqualFold(act.ResourceID, resourceID) && (strings.EqualFold(act.ProjectID, projectID) || len(projectID) == 0) {
			ram.WeekActivities[i].HoursAvailableForProject -= hoursToSubtract
		}
	}
}

// GetDistinctWeeksInFlatmap return a list of CSWeeks from the flatmap
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

// FlattenPortfolio takes the elements of ths scheduled portfolio and flattens them into a tabular format
func FlattenPortfolio(scheduleList []Schedule, rm map[string]resource.Resource) ([]ResourceProjectHourAllocation, error) {
	out := []ResourceProjectHourAllocation{}

	for _, proj := range scheduleList {
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
