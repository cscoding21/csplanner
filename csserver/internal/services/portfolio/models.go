package portfolio

import (
	"csserver/internal/services/schedule"
	"time"
)

type Portfolio struct {
	Schedule []schedule.Schedule
}

func (p *Portfolio) GetDateRange() (*time.Time, *time.Time) {
	start := time.Now().AddDate(10, 0, 0)
	end := start.AddDate(-10, 0, 0)

	if len(p.Schedule) == 0 {
		return nil, nil
	}

	for _, s := range p.Schedule {
		if len(s.ProjectActivityWeeks) == 0 {
			continue
		}

		if s.ProjectActivityWeeks[0].Begin.Before(start) {
			start = s.ProjectActivityWeeks[0].Begin
		}

		if s.ProjectActivityWeeks[len(s.ProjectActivityWeeks)-1].End.After(end) {
			end = s.ProjectActivityWeeks[len(s.ProjectActivityWeeks)-1].End
		}
	}

	return &start, &end
}
