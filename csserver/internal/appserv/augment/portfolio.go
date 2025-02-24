package augment

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/calendar"
	"csserver/internal/common"
	"csserver/internal/services/resource/rtypes/resourcestatus"
	"csserver/internal/services/resource/rtypes/resourcetype"
)

// AugmentPortfolio enhance calculate and caches properties of a portfolio
func AugmentPortfolio(port *idl.Portfolio, resourceID *string) {
	//---exit conditions
	if port.Begin == nil || port.End == nil {
		return
	}

	resourceList := findResources()

	orgCapacity := 0

	for _, r := range *resourceList {
		if resourceID != nil && r.ID != *resourceID {
			continue
		}

		if resourceID != nil && r.Type == resourcetype.Human {
			orgCapacity += r.AvailableHoursPerWeek
		} else if r.Type == resourcetype.Human && r.Status == resourcestatus.Inhouse {
			orgCapacity += r.AvailableHoursPerWeek
		}
	}

	weeks := calendar.GetWeeks(*port.Begin, *port.End)

	portWeeks := []*idl.PortfolioWeekSummary{}

	for _, w := range weeks {
		thisWeek := idl.PortfolioWeekSummary{
			WeekNumber:     w.WeekNumber,
			Year:           w.Year,
			Begin:          w.Begin,
			End:            w.End,
			OrgCapacity:    orgCapacity,
			AllocatedHours: getWeekHourSummary(common.RefToValSlice(port.Schedule), w),
		}

		portWeeks = append(portWeeks, &thisWeek)
	}

	port.WeekSummary = portWeeks
	calculated := getPortfolioCalculatedData(port)

	port.Calculated = &calculated
}

func getWeekHourSummary(sch []idl.Schedule, week calendar.CSWeek) int {
	out := 0
	for _, s := range sch {
		for _, w := range s.ProjectActivityWeeks {
			if w.End.Equal(week.End) {
				for _, a := range w.Activities {
					out += a.HoursSpent
				}
			}
		}
	}

	return out
}

func getPortfolioCalculatedData(port *idl.Portfolio) idl.PortfolioCalculatedData {
	out := idl.PortfolioCalculatedData{
		ValueInFlight:  0.0,
		ValueScheduled: 0.0,
		TotalValue:     0.0,
		CountInFlight:  0,
		CountScheduled: 0,
		TotalCount:     0,
	}

	for _, s := range port.Schedule {
		out.TotalCount++
		out.TotalValue += *s.Project.ProjectValue.Calculated.NetPresentValue
		if s.Project.ProjectStatusBlock.Status == "inflight" {
			out.CountInFlight++
			out.ValueInFlight += *s.Project.ProjectValue.Calculated.NetPresentValue
		} else if s.Project.ProjectStatusBlock.Status == "scheduled" {
			out.CountScheduled++
			out.ValueScheduled += *s.Project.ProjectValue.Calculated.NetPresentValue
		}
	}

	return out
}
