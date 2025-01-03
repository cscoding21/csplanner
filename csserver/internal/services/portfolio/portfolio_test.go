package portfolio

import (
	"csserver/internal/calendar"
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

func getTestPortfolio() Portfolio {
	p := Portfolio{}

	sc := ProjectSchedule{ProjectID: "project:test", ProjectName: "test"}

	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
	weeks := calendar.GetWeeks(start, end)

	for _, w := range weeks {
		sc.ProjectActivityWeeks = append(sc.ProjectActivityWeeks, &ProjectActivityWeek{
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
		sc.ProjectActivityWeeks = append(sc.ProjectActivityWeeks, &ProjectActivityWeek{
			WeekNumber: w.WeekNumber,
			Year:       w.Year,
			Begin:      w.Begin,
			End:        w.End,
		})
	}

	p.Schedule = append(p.Schedule, sc)

	return p
}
