package augment

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/calendar"
	"csserver/internal/common"
	"csserver/internal/services/resource/rtypes/resourcestatus"
	"csserver/internal/services/resource/rtypes/resourcetype"
)

// AugmentPortfolio enhance calculate and caches properties of a portfolio
func AugmentPortfolio(port *idl.Portfolio) {
	//---exit conditions
	if port.Begin == nil || port.End == nil {
		return
	}

	resourceList := findResources()

	orgCapacity := 0

	for _, r := range *resourceList {
		if r.Type == resourcetype.Human && r.Status == resourcestatus.Inhouse {
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
