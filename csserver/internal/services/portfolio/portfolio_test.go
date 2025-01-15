package portfolio_test

import (
	"csserver/internal/calendar"
	"csserver/internal/services/portfolio"
	"csserver/internal/services/schedule"
	"fmt"
	"testing"
	"time"
)

func TestGetDateRange(t *testing.T) {
	p := getTestPortfolio()

	if len(p.Schedule) == 0 {
		t.Error("test schedule did not load correctly")
	}

	start, end := p.GetDateRange()

	fmt.Println(start)
	fmt.Println(end)
}

func getTestPortfolio() portfolio.Portfolio {
	p := portfolio.Portfolio{}

	sc := schedule.Schedule{ProjectID: "project:test", ProjectName: "test"}

	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
	weeks := calendar.GetWeeks(start, end)

	for _, w := range weeks {
		sc.ProjectActivityWeeks = append(sc.ProjectActivityWeeks, &schedule.ProjectActivityWeek{
			WeekNumber: w.WeekNumber,
			Year:       w.Year,
			Begin:      w.Begin,
			End:        w.End,
		})
	}

	p.Schedule = append(p.Schedule, sc)

	start = time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)
	end = time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC)
	weeks = calendar.GetWeeks(start, end)

	for _, w := range weeks {
		sc.ProjectActivityWeeks = append(sc.ProjectActivityWeeks, &schedule.ProjectActivityWeek{
			WeekNumber: w.WeekNumber,
			Year:       w.Year,
			Begin:      w.Begin,
			End:        w.End,
		})
	}

	p.Schedule = append(p.Schedule, sc)

	return p
}
