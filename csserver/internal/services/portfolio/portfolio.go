package portfolio

import (
	"csserver/internal/config"
	"csserver/internal/interfaces"
	"csserver/internal/providers/nats"
	"csserver/internal/providers/surreal"
)

// ---This is the name of the object in the database
const PortfolioIdentifier = "portfolio"

// OrganizationService is a service for interacting with lists.
type PortfolioService struct {
	DBClient      surreal.DBClient
	ContextHelper interfaces.ContextHelpers
	PubSub        nats.PubSubProvider
}

// NewOrganizationService creates a new Organization service.
func NewPortfolioService(
	db surreal.DBClient,
	ch config.ContextHelper,
	ps nats.PubSubProvider) *PortfolioService {

	return &PortfolioService{
		DBClient:      db,
		ContextHelper: &ch,
		PubSub:        ps,
	}
}

func (service *PortfolioService) LoadResourceUtilization() int {
	return 0
}
