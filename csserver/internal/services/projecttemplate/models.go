//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package projecttemplate

import (
	"csserver/internal/common"
)

type Projecttemplate struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Phases      []ProjecttemplatePhase `json:"phases"`
}

type ProjecttemplatePhase struct {
	ID          string                `json:"id,omitempty"`
	PhaseOrder  byte                  `json:"phase_order"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Tasks       []ProjectTemplateTask `json:"tasks"`
}

type ProjectTemplateTask struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	RequiredSkillID string `json:"required_skill_id"`
}
