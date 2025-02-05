package processor

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/interfaces"
	"csserver/internal/providers/nats"
	"csserver/internal/providers/surreal"
	"csserver/internal/services/project"
	"csserver/internal/services/project/ptypes/milestonestatus"
	"csserver/internal/services/project/ptypes/projectstatus"
	"csserver/internal/services/resource"
	"csserver/internal/services/schedule"
	"errors"
	"time"
)

// ---This is the name of the object in the database
const ProcessorIdentifier = "autoprocessor"

// PortfolioService is a service for interacting with the project portfolio.
type ProcessorService struct {
	DBClient        surreal.DBClient
	ContextHelper   interfaces.ContextHelpers
	PubSub          nats.PubSubProvider
	ProjectService  project.ProjectService
	ResourceService resource.ResourceService
	ScheduleService schedule.ScheduleService
}

// NewProcessorService creates a new portfolio service.
func NewProcessorService(
	db surreal.DBClient,
	ch config.ContextHelper,
	ps nats.PubSubProvider,
	projService project.ProjectService,
	resService resource.ResourceService,
	schedService schedule.ScheduleService) *ProcessorService {

	return &ProcessorService{
		DBClient:        db,
		ContextHelper:   &ch,
		PubSub:          ps,
		ProjectService:  projService,
		ResourceService: resService,
		ScheduleService: schedService,
	}
}

func (s *ProcessorService) ProcessNightly(ctx context.Context) error {
	projects, err := s.ProjectService.FindAllProjects(ctx)
	if err != nil {
		return err
	}

	err = s.ProcessSetProjectToInFlight(ctx, projects.Results)
	if err != nil {
		return err
	}

	err = s.ProcessCompleteProject(ctx, projects.Results)
	if err != nil {
		return err
	}

	return nil
}

// ProcessSetProjectToInFlight Check the status of a scheduled project to determine if it is in flight
func (s *ProcessorService) ProcessSetProjectToInFlight(ctx context.Context, projects []project.Project) error {
	var e error

	for _, p := range projects {
		if p.ProjectBasics == nil || p.ProjectStatusBlock == nil || p.ProjectBasics.StartDate == nil {
			continue
		}
		if p.ProjectStatusBlock.Status == projectstatus.Scheduled && p.ProjectBasics.StartDate.Before(time.Now()) {
			_, err := s.ProjectService.SetProjectStatus(ctx, p.ID, projectstatus.InFlight, true)
			if err != nil {
				errors.Join(e, err)
			}
		}
	}

	return e
}

// ProcessCompleteProject check the status of all project tasks to see if the project is complete
func (s *ProcessorService) ProcessCompleteProject(ctx context.Context, projects []project.Project) error {
	var e error

	for _, p := range projects {
		if p.ProjectStatusBlock.Status == projectstatus.Complete {
			continue
		}

		incompleteCount := 0
		for _, m := range p.ProjectMilestones {
			for _, t := range m.Tasks {
				if !common.IsOneOf(t.Status, milestonestatus.Done, milestonestatus.Removed) {
					incompleteCount++
				}
			}
		}

		if incompleteCount == 0 {
			_, err := s.ProjectService.SetProjectStatus(ctx, p.ID, projectstatus.Complete, true)
			if err != nil {
				errors.Join(e, err)
			}
		}

	}

	return e
}
