package project

import (
	"csserver/internal/services/project/ptypes/projectstatus"
	"testing"
)

func TestSetProjectState(t *testing.T) {
	testData := []struct {
		P        Project
		ToState  projectstatus.ProjectState
		Expected bool
	}{
		//---From New
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.NewProject, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.Draft, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.Proposed, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.Approved, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.Rejected, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.Backlogged, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.Scheduled, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.InFlight, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.Complete, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.Deferred, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.NewProject}}, ToState: projectstatus.Abandoned, Expected: false},

		//---From Draft
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.Draft, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.Proposed, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.Approved, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.Rejected, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.Backlogged, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.Scheduled, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.InFlight, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.Complete, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.Deferred, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Draft}}, ToState: projectstatus.Abandoned, Expected: false},

		//---From Proposed
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.Draft, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.Proposed, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.Approved, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.Rejected, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.Backlogged, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.Scheduled, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.InFlight, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.Complete, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.Deferred, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Proposed}}, ToState: projectstatus.Abandoned, Expected: false},

		//---From Approved
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.Draft, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.Proposed, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.Approved, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.Rejected, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.Backlogged, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.Scheduled, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.InFlight, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.Complete, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.Deferred, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Approved}}, ToState: projectstatus.Abandoned, Expected: false},

		//---From Rejected
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.Draft, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.Proposed, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.Approved, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.Rejected, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.Backlogged, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.Scheduled, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.InFlight, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.Complete, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.Deferred, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Rejected}}, ToState: projectstatus.Abandoned, Expected: false},

		//---From Backlogged
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.Draft, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.Proposed, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.Approved, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.Rejected, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.Backlogged, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.Scheduled, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.InFlight, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.Complete, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.Deferred, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Backlogged}}, ToState: projectstatus.Abandoned, Expected: false},

		//---From Scheduled
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.Draft, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.Proposed, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.Approved, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.Rejected, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.Backlogged, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.Scheduled, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.InFlight, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.Complete, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.Deferred, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Scheduled}}, ToState: projectstatus.Abandoned, Expected: true},

		//---From InFlight
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.Draft, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.Proposed, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.Approved, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.Rejected, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.Backlogged, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.Scheduled, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.InFlight, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.Complete, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.Deferred, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.InFlight}}, ToState: projectstatus.Abandoned, Expected: false},

		//---From Complete
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.Draft, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.Proposed, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.Approved, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.Rejected, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.Backlogged, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.Scheduled, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.InFlight, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.Complete, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.Deferred, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Complete}}, ToState: projectstatus.Abandoned, Expected: false},

		//---From Deferred
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.Draft, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.Proposed, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.Approved, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.Rejected, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.Backlogged, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.Scheduled, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.InFlight, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.Complete, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.Deferred, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Deferred}}, ToState: projectstatus.Abandoned, Expected: false},

		//---From Abandoned
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.NewProject, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.Draft, Expected: true},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.Proposed, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.Approved, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.Rejected, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.Backlogged, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.Scheduled, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.InFlight, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.Complete, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.Deferred, Expected: false},
		{P: Project{ProjectBasics: &ProjectBasics{Status: projectstatus.Abandoned}}, ToState: projectstatus.Abandoned, Expected: true},
	}

	for _, td := range testData {
		newState := stateMachineMap[td.ToState]

		result := newState.Can(&td.P)
		if td.Expected != result.Pass {
			t.Error(result.Messages)
		}
	}
}
