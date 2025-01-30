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
	State           projectstatus.ProjectState
	Description     string
	NextValidStates []projectstatus.ProjectState

	Can func(p *Project) validate.ValidationResult
}

var stateMachineMap = map[projectstatus.ProjectState]ProjectStatus{
	//---COMPLETE
	projectstatus.NewProject: {
		State:       projectstatus.NewProject,
		Description: "the state when project first created",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Draft},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()

			// this state can only be entered on initial project creation
			// it cannot be transioned at any other time

			return result
		},
	},

	projectstatus.Draft: {
		State:       projectstatus.Draft,
		Description: "the state when project details are being actively formulated",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Proposed, projectstatus.Abandoned},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.Draft

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.NewProject,
				projectstatus.Draft,
				projectstatus.Proposed,
				projectstatus.Approved,
				projectstatus.Rejected,
				projectstatus.Backlogged,
				projectstatus.Scheduled,
				projectstatus.InFlight,
				projectstatus.Deferred,
				projectstatus.Abandoned,
				projectstatus.Exception))

			return result
		},
	},

	projectstatus.Proposed: {
		State:       projectstatus.Proposed,
		Description: "project details are complete and is ready for review",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Approved, projectstatus.Rejected, projectstatus.Abandoned},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.Proposed

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Draft))

			return result
		},
	},

	projectstatus.Approved: {
		State:       projectstatus.Approved,
		Description: "project has been reviewed and approved",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Scheduled, projectstatus.Backlogged, projectstatus.Draft},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.Approved

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Proposed))

			return result
		},
	},

	projectstatus.Rejected: {
		State:       projectstatus.Rejected,
		Description: "project has been reviewed and rejected",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Abandoned, projectstatus.Draft},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.Rejected

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Proposed))

			return result
		},
	},

	projectstatus.Backlogged: {
		State:       projectstatus.Backlogged,
		Description: "project is approved and awaiting resource availability",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Scheduled, projectstatus.Draft, projectstatus.Abandoned},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.Backlogged

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Approved))

			return result
		},
	},

	projectstatus.Scheduled: {
		State:       projectstatus.Scheduled,
		Description: "project has been staffed and a start date has been set",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Deferred},
		//---can auto-transition to in-flight

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.Scheduled

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Approved,
				projectstatus.Deferred,
				projectstatus.Backlogged))

			if p.ProjectBasics.StartDate == nil {
				result.Append(getFailingMessage("Start date", "Project start date must be set to schedule implementation work"))
			}

			return result
		},
	},

	projectstatus.InFlight: {
		State:       projectstatus.InFlight,
		Description: "project has been started",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Deferred, projectstatus.Abandoned},
		//---can auto-transition to complete

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.InFlight

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Deferred))
			//---can auto-transition to scheduled if start date is set to the future

			return result
		},
	},

	projectstatus.Complete: {
		State:       projectstatus.Complete,
		Description: "project has been completed",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.Complete

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.InFlight))

			return result
		},
	},

	projectstatus.Deferred: {
		State:       projectstatus.Deferred,
		Description: "project was in flight, but work has been postponed",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.InFlight, projectstatus.Scheduled},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.Deferred

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.Draft,
				projectstatus.Abandoned,
				projectstatus.Scheduled))

			return result
		},
	},

	projectstatus.Abandoned: {
		State:       projectstatus.Abandoned,
		Description: "project made it past scheduled, but has been cancelled",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Draft},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()
			currentStatus := p.ProjectStatusBlock.Status
			thisState := projectstatus.Abandoned

			//---states that can be transitioned into this state
			result.Append(checkStateTransition(currentStatus, thisState,
				projectstatus.NewProject,
				projectstatus.Draft,
				projectstatus.Proposed,
				projectstatus.Approved,
				projectstatus.Rejected,
				projectstatus.Backlogged,
				projectstatus.Deferred))

			return result
		},
	},

	projectstatus.Exception: {
		State:       projectstatus.Exception,
		Description: "denotes a mapping error...should never happen",
		//---states that can be entered from this state
		NextValidStates: []projectstatus.ProjectState{projectstatus.Draft},

		Can: func(p *Project) validate.ValidationResult {
			result := validate.NewSuccessValidationResult()

			//---states that can be transitioned into this state
			// this state can only be entered by forcing a bypass
			// it cannot occur organically

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

func getFailingMessage(field, message string) validate.ValidationResult {
	vm := validate.NewValidationMessage(field, message)
	out := validate.NewFailingValidationResult(vm)

	return out
}
