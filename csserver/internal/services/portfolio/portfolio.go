package portfolio

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/interfaces"
	"csserver/internal/providers/nats"
	"csserver/internal/providers/surreal"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"csserver/internal/services/schedule"
	"fmt"
	"time"
)

// ---This is the name of the object in the database
const PortfolioIdentifier = "portfolio"

// PortfolioService is a service for interacting with the project portfolio.
type PortfolioService struct {
	DBClient        surreal.DBClient
	ContextHelper   interfaces.ContextHelpers
	PubSub          nats.PubSubProvider
	ProjectService  project.ProjectService
	ResourceService resource.ResourceService
	ScheduleService schedule.ScheduleService
}

// NewPortfolioService creates a new portfolio service.
func NewPortfolioService(
	db surreal.DBClient,
	ch config.ContextHelper,
	ps nats.PubSubProvider,
	projService project.ProjectService,
	resService resource.ResourceService,
	schedService schedule.ScheduleService) *PortfolioService {

	return &PortfolioService{
		DBClient:        db,
		ContextHelper:   &ch,
		PubSub:          ps,
		ProjectService:  projService,
		ResourceService: resService,
		ScheduleService: schedService,
	}
}

// ScheduleProject add a project to the current portfolio
func (ps *PortfolioService) ScheduleProject(ctx context.Context, projectID string, startDate time.Time) (Portfolio, error) {
	panic("not implemented")
}

// GetUnbalancedPortfolio retrieves the currently scheduled projects
func (ps *PortfolioService) GetUnbalancedPortfolio(ctx context.Context) (*Portfolio, error) {
	pf := common.NewPagedResultsForAllRecords[project.Project]()
	pf.Filters.AddFilter(common.Filter{Key: "basics.status", Value: "scheduled,inflight", Operation: common.FilterOperationIn})

	rm, err := ps.ResourceService.GetResourceMap(ctx, false)
	if err != nil {
		return nil, err
	}

	ram, err := ps.ScheduleService.GetInitialResourceAllocationMap(rm)
	if err != nil {
		return nil, err
	}

	response, err := ps.ProjectService.FindProjects(ctx, pf.Pagination, pf.Filters)
	if err != nil {
		return nil, err
	}

	port := NewPortfolio(response.Results, rm)

	for _, proj := range response.Results {
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

	return &port, nil
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

	port.Validator = &comparer

	return nil
}
