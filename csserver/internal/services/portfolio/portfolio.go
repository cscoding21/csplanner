package portfolio

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/events"
	"csserver/internal/interfaces"
	"csserver/internal/providers/postgres"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"csserver/internal/services/schedule"
	"csserver/internal/utils"
	"fmt"
	"sort"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ---This is the name of the object in the database
const PortfolioIdentifier = "portfolio"

// PortfolioService is a service for interacting with the project portfolio.
type PortfolioService struct {
	DBClient        *pgxpool.Pool
	ContextHelper   interfaces.ContextHelpers
	PubSub          events.PubSubProvider
	ProjectService  project.ProjectService
	ResourceService resource.ResourceService
	ScheduleService schedule.ScheduleService
}

// NewPortfolioService creates a new portfolio service.
func NewPortfolioService(
	db *pgxpool.Pool,
	ps events.PubSubProvider,
	projService project.ProjectService,
	resService resource.ResourceService,
	schedService schedule.ScheduleService) *PortfolioService {

	return &PortfolioService{
		DBClient:        db,
		PubSub:          ps,
		ProjectService:  projService,
		ResourceService: resService,
		ScheduleService: schedService,
	}
}

// GetUnbalancedPortfolio retrieves the currently scheduled projects
func (ps *PortfolioService) GetUnbalancedPortfolio(ctx context.Context, additionalProjectID string) (*Portfolio, error) {
	pf := common.NewPagedResultsForAllRecords[project.Project]()

	pf.Filters.AddFilter(common.Filter{Key: "status1", Value: "scheduled", Operation: common.FilterOperationIn})
	pf.Filters.AddFilter(common.Filter{Key: "status2", Value: "inflight", Operation: common.FilterOperationIn})

	pf.Filters.AddFilter(common.Filter{Key: "additionalID", Value: additionalProjectID, Operation: common.FilterOperationEqual})

	rm, err := ps.ResourceService.GetResourceMap(ctx, false)
	if err != nil {
		return nil, err
	}

	ram, err := ps.ScheduleService.GetInitialResourceAllocationMap(rm)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(`
		SELECT * 
		FROM %s 
		WHERE true 
			and (
				data->'status'->>'status' IN ($1, $2) 
				or id = $3
		)
		ORDER BY data->'basics'->>'start_time'`, project.ProjectIdentifier)
	response, err := postgres.FindPagedObjects[project.Project](ctx, ps.DBClient, sql, pf.Pagination, pf.Filters, pf.Filters.GetFiltersOrderedValues())
	if err != nil {
		return nil, err
	}

	port := NewPortfolio(common.ExtractDataFromBase(response.Results), rm)

	for _, pr := range response.Results {
		proj := pr.Data
		startDate := time.Now()
		if proj.ProjectBasics.StartDate != nil {
			startDate = *proj.ProjectBasics.StartDate
		}

		sch, err := ps.ScheduleService.CalculateProjectSchedule(ctx, &proj, startDate, ram)
		if err != nil {
			return &port, err
		}

		port.Schedule = append(port.Schedule, sch)
	}

	sort.Slice(port.Schedule, func(p, q int) bool {
		return port.Schedule[p].End.Before(port.Schedule[q].End)
	})

	return &port, nil
}

// GetBalancedPortfolio returns a portfolio that has been run through the balancing process
func (ps *PortfolioService) GetBalancedPortfolio(ctx context.Context, additionalProjectID string) (*Portfolio, error) {
	ubp, err := ps.GetUnbalancedPortfolio(ctx, additionalProjectID)
	if err != nil {
		return nil, nil
	}

	err = ps.BalancePortfolio(ctx, ubp)
	if err != nil {
		return nil, nil
	}

	sort.Slice(ubp.Schedule, func(p, q int) bool {
		return ubp.Schedule[p].End.Before(ubp.Schedule[q].End)
	})

	return ubp, nil

}

// BalancePortfolio iterate over the portfolio until it is balanced with proper resource utilization
func (ps *PortfolioService) BalancePortfolio(ctx context.Context, port *Portfolio) error {
	comparer := NewPortfolioComparer()
	maxIterationCount := 256
	iterationCount := 0

	for !comparer.Validate(&port.Schedule) {
		if iterationCount > maxIterationCount {
			return fmt.Errorf("portfolio balance process exceeded max iteration count")
		}

		ram, err := ps.ScheduleService.GetResourceAllocationMap(port.Schedule, port.ResourceMap)
		if err != nil {
			return err
		}

		for i, s := range port.Schedule {
			proj := port.ProjectMap[s.ProjectID]
			port.Schedule[i], err = ps.ScheduleService.CalculateProjectSchedule(ctx, &proj, *proj.ProjectBasics.StartDate, ram)
			if err != nil {
				return err
			}
		}
	}

	err := port.SetCalculatedData()
	if err != nil {
		return err
	}

	port.Validator = &comparer

	return nil
}

// GetPortfolioForResource retrieve a portfolio only containing tasks relative to a user
func (ps *PortfolioService) GetPortfolioForResource(ctx context.Context, port *Portfolio, resourceID string) (*Portfolio, error) {
	outPort := Portfolio{
		ResourceMap: make(map[string]resource.Resource),
		ProjectMap:  port.ProjectMap,
		Schedule:    []schedule.Schedule{},
	}

	//---copy only the in-scope resource to the map
	outPort.ResourceMap[resourceID] = port.ResourceMap[resourceID]

	for _, sch := range port.Schedule {
		addSchedule := false

		newSch := schedule.Schedule{
			ProjectID:            sch.ProjectID,
			ProjectName:          sch.ProjectName,
			Begin:                sch.Begin,
			End:                  sch.End,
			ProjectActivityWeeks: []*schedule.ProjectActivityWeek{},
		}
		for _, week := range sch.ProjectActivityWeeks {
			newWeek := schedule.ProjectActivityWeek{
				Begin:      week.Begin,
				End:        week.End,
				WeekNumber: week.WeekNumber,
				Year:       week.Year,
				Activities: []schedule.ProjectActivity{},
			}

			newSch.ProjectActivityWeeks = append(newSch.ProjectActivityWeeks, &newWeek)
			for _, act := range week.Activities {
				if act.ResourceID == resourceID {
					addSchedule = true

					newAct, err := utils.DeepCopy(act)
					if err != nil {
						return nil, err
					}

					newWeek.Activities = append(newWeek.Activities, *newAct)
				}
			}
		}

		if addSchedule {
			outPort.Schedule = append(outPort.Schedule, newSch)
		}
	}

	return &outPort, nil
}
