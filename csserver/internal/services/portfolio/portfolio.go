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

// GetCurrentPortfolio retrieves the currently scheduled projects
func (ps *PortfolioService) GetCurrentPortfolio(ctx context.Context, ram schedule.ResourceAllocationMap) (Portfolio, error) {
	port := Portfolio{}
	compareMap := PortfolioComparer{}

	pf := common.NewPagedResultsForAllRecords[project.Project]()
	pf.Filters.AddFilter(common.Filter{Key: "basics.status", Value: "scheduled,inflight", Operation: common.FilterOperationIn})

	response, err := ps.ProjectService.FindProjects(ctx, pf.Pagination, pf.Filters)
	if err != nil {
		return port, err
	}

	for _, proj := range response.Results {
		startDate := time.Now()
		if proj.ProjectBasics.StartDate != nil {
			startDate = *proj.ProjectBasics.StartDate
		}

		sch, err := ps.ScheduleService.CalculateProjectSchedule(ctx, &proj, startDate, ram)
		if err != nil {
			return port, err
		}

		port.Schedule = append(port.Schedule, sch)
	}

	fmt.Println(compareMap)
	return port, nil
}

func iteratePortfolio(scheduleList []schedule.Schedule) {

}
