//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package projecttemplate

import (
	"csserver/internal/common"
	"csserver/internal/services/project/ptypes"
)

type Projecttemplate struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	Name   string                 `json:"name"`
	Phases []ProjecttemplatePhase `json:"phases"`
}

type ProjecttemplatePhase struct {
	ID          string                 `json:"id,omitempty"`
	PhaseOrder  byte                   `json:"phase_order"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Tasks       []ProjectTemplateTasks `json:"tasks"`
}

type ProjectTemplateTasks struct {
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	Status           ptypes.MilestoneStatus `json:"status"`
	RequiredSkillIDs []string               `json:"required_skill_ids"`
}
