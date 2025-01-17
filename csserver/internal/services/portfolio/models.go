package portfolio

import (
	"csserver/internal/services/schedule"
	"time"
)

type PortfolioComparer struct {
	CompareMap map[string]HashResult
}

type HashResult struct {
	LastHash       uint64
	CurrentHash    uint64
	IterationCount int
}

func (pc *PortfolioComparer) Validate(scheduleList []schedule.Schedule) bool {
	out := true
	for _, s := range scheduleList {
		item, ok := pc.CompareMap[s.ProjectID]
		if !ok {
			out = false
			pc.CompareMap[s.ProjectID] = HashResult{
				LastHash:       0,
				CurrentHash:    s.Hash,
				IterationCount: 1,
			}

			continue
		} else {
			if s.Hash == item.CurrentHash {
				//---this one is OK
				continue
			} else {
				out = false
			}

			newItem := HashResult{
				LastHash:       item.CurrentHash,
				CurrentHash:    s.Hash,
				IterationCount: item.IterationCount + 1,
			}

			pc.CompareMap[s.ProjectID] = newItem
		}
	}

	return out
}

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
