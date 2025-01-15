package portfolio

import (
	"csserver/internal/config"
	"csserver/internal/interfaces"
	"csserver/internal/providers/nats"
	"csserver/internal/providers/surreal"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
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
}

// NewPortfolioService creates a new portfolio service.
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

// GetResourceUtilizationTable get an enhanced resource table that includes project allocations
// func (service *PortfolioService) GetResourceUtilizationTable(ctx context.Context) (*ResourceUtilizationTable, error) {
// 	out := ResourceUtilizationTable{
// 		Resources: []ResourceUtilizationItem{},
// 	}

// 	pf := common.NewPagedResultsForAllRecords[project.Project]()
// 	pf.Filters.AddFilter(common.Filter{
// 		Key:       "basics.status",
// 		Value:     "draft,",
// 		Operation: "in",
// 	})

// 	rm, _ := service.ResourceService.GetResourceMap(ctx, false)
// 	projects, err := service.ProjectService.FindProjects(ctx, pf.Pagination, pf.Filters)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, p := range projects.Results {
// 		for _, m := range p.ProjectMilestones {
// 			for _, t := range m.Tasks {
// 				for _, resID := range t.ResourceIDs {
// 					thisResource := rm[resID]
// 					//---project info
// 					r := ResourceUtilizationItem{
// 						ProjectID:     p.ID,
// 						ProjectName:   p.ProjectBasics.Name,
// 						ProjectStatus: p.ProjectBasics.Status,
// 						ResourceID:    resID,
// 						ResourceName:  thisResource.Name,
// 					}

// 					//---milestone info
// 					r.MilestoneName = m.Phase.Name

// 					//---task info
// 					r.TaskName = t.Name
// 					r.TaskHourEstimate = t.HourEstimate

// 					out.Resources = append(out.Resources, r)
// 				}
// 			}
// 		}
// 	}

// 	return &out, nil
// }
