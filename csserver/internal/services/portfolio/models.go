package portfolio

import (
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"csserver/internal/services/schedule"
	"time"
)

func NewPortfolioComparer() PortfolioComparer {
	pc := PortfolioComparer{CompareMap: make(map[string]HashResult)}

	return pc
}

type PortfolioComparer struct {
	Iterations int
	CompareMap map[string]HashResult
}

type HashResult struct {
	LastHash       uint64
	CurrentHash    uint64
	IterationCount int
}

func (pc *PortfolioComparer) Validate(scheduleList *[]schedule.Schedule) bool {
	pc.Iterations++
	out := true
	for _, s := range *scheduleList {
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
				pc.CompareMap[s.ProjectID] = HashResult{
					LastHash:       item.CurrentHash,
					CurrentHash:    s.Hash,
					IterationCount: item.IterationCount,
				}

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

func NewPortfolio(projectList []project.Project, rm map[string]resource.Resource) Portfolio {
	port := Portfolio{ProjectMap: make(map[string]project.Project), ResourceMap: rm}

	for _, p := range projectList {
		port.ProjectMap[p.ID] = p
	}

	return port
}

type Portfolio struct {
	ProjectMap  map[string]project.Project
	ResourceMap map[string]resource.Resource
	Schedule    []schedule.Schedule
	Validator   *PortfolioComparer
}

type PortfolioCalculatedData struct {
	CountInFlight  int
	ValueInFlight  float64
	CountScheduled int
	ValueScheduled float64

	AggregatedNPV float64
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
