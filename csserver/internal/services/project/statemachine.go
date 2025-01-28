package project

import (
	"csserver/internal/common"
	"csserver/internal/services/project/ptypes/projectstatus"
	"fmt"

	"github.com/cscoding21/csval/validate"
)

/*
ProjectStateNotSet ProjectState = "notset"

Exception  ProjectState = "exception"  //---denotes a mapping error...should never happen
NewProject ProjectState = "new"        //---when project first created
Draft      ProjectState = "draft"      //---project details (features, cost, value) are being fleshed out
Proposed   ProjectState = "proposed"   //---project details are complete and is ready for review
Approved   ProjectState = "approved"   //---project has been reviewed and approved
Rejected   ProjectState = "rejected"   //---project has been reviewed and rejected
Backlogged ProjectState = "backlogged" //---project is approved and awaiting resource availability
Scheduled  ProjectState = "scheduled"  //---project has been staffed and a start date has been set
InFlight   ProjectState = "inflight"   //---project has been started
Complete   ProjectState = "complete"   //---project has been completed
Deferred   ProjectState = "deferred"   //---project was in flight, but work has been postponed
Abandoned  ProjectState = "abandoned"  //---project made it past scheduled, but has been cancelled
*/

type ProjectStatus struct {
	State       projectstatus.ProjectState
	Description string

	Can func(p *Project) validate.ValidationResult
}

type StatusCondition struct {
	Description string
	Met         bool
}

var stateMachineMap = map[projectstatus.ProjectState]ProjectStatus{
	projectstatus.NewProject: {
		State:       projectstatus.NewProject,
		Description: "the state when project first created",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.NewProject

			result.Append(checkStateTransition(currentStatus, thisState, projectstatus.NewProject))

			return result
		},
	},
	projectstatus.Draft: {
		State:       projectstatus.Draft,
		Description: "the state when project details are being actively formulated",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.Draft

			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.NewProject,
				projectstatus.Proposed,
				projectstatus.Approved,
				projectstatus.Rejected,
				projectstatus.Backlogged,
				projectstatus.Deferred,
				projectstatus.Abandoned))

			return result
		},
	},

	projectstatus.Proposed: {
		State:       projectstatus.Proposed,
		Description: "project details are complete and is ready for review",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.Proposed

			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Draft))

			return result
		},
	},

	projectstatus.Approved: {
		State:       projectstatus.Approved,
		Description: "project has been reviewed and approved",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.Approved

			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Proposed))

			return result
		},
	},

	projectstatus.Rejected: {
		State:       projectstatus.Rejected,
		Description: "project has been reviewed and rejected",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.Rejected

			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Proposed,
				projectstatus.Approved,
				projectstatus.Backlogged))

			return result
		},
	},

	projectstatus.Backlogged: {
		State:       projectstatus.Backlogged,
		Description: "project is approved and awaiting resource availability",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.Backlogged

			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Approved))

			return result
		},
	},

	projectstatus.Scheduled: {
		State:       projectstatus.Scheduled,
		Description: "project has been staffed and a start date has been set",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.Scheduled

			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Approved))

			return result
		},
	},

	projectstatus.InFlight: {
		State:       projectstatus.InFlight,
		Description: "project has been started",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.InFlight

			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Scheduled))

			return result
		},
	},

	projectstatus.Complete: {
		State:       projectstatus.Complete,
		Description: "project has been completed",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.Complete

			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.InFlight))

			return result
		},
	},

	projectstatus.Deferred: {
		State:       projectstatus.Deferred,
		Description: "project was in flight, but work has been postponed",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.Deferred

			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.InFlight))

			return result
		},
	},

	projectstatus.Abandoned: {
		State:       projectstatus.Abandoned,
		Description: "project made it past scheduled, but has been cancelled",

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectBasics.Status
			thisState := projectstatus.Abandoned

			result.Append(checkStateTransition(currentStatus, thisState,
				thisState,
				projectstatus.Scheduled))

			return result
		},
	},
}

func checkStateTransition(currentState projectstatus.ProjectState, proposedState projectstatus.ProjectState, approvedStates ...projectstatus.ProjectState) validate.ValidationResult {
	if currentState == proposedState || common.IsOneOf(currentState, approvedStates...) {
		return validate.NewSuccessValidationResult()
	}

	msg := validate.NewValidationMessage("status", fmt.Sprintf("Project cannot move from state '%v' to '%v'", currentState, proposedState))

	return validate.NewFailingValidationResult(msg)
}
