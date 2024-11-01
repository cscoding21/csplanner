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
	"fmt"
)

// ---This is the name of the object in the database
const PortfolioIdentifier = "portfolio"

// OrganizationService is a service for interacting with lists.
type PortfolioService struct {
	DBClient        surreal.DBClient
	ContextHelper   interfaces.ContextHelpers
	PubSub          nats.PubSubProvider
	ProjectService  project.ProjectService
	ResourceService resource.ResourceService
}

// NewOrganizationService creates a new Organization service.
func NewPortfolioService(
	db surreal.DBClient,
	ch config.ContextHelper,
	ps nats.PubSubProvider,
	projService project.ProjectService,
	resService resource.ResourceService) *PortfolioService {

	return &PortfolioService{
		DBClient:        db,
		ContextHelper:   &ch,
		PubSub:          ps,
		ProjectService:  projService,
		ResourceService: resService,
	}
}

func (service *PortfolioService) GetResourceUtilizationTable(ctx context.Context) (*ResourceUtilizationTable, error) {
	out := ResourceUtilizationTable{}

	pf := common.NewPagedResultsForAllRecords[project.Project]()
	pf.Filters.AddFilter(common.Filter{
		Key:       "basics.status",
		Value:     "",
		Operation: "lk",
	})

	projects, err := service.ProjectService.FindProjects(ctx, pf.Pagination, pf.Filters)
	if err != nil {
		return nil, err
	}

	for _, p := range projects.Results {
		fmt.Println(p)
	}

	return &out, nil
}
