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
	Name        string                  `json:"name" csval:"req"`
	Description string                  `json:"description" csval:"req"`
	Phases      []*ProjecttemplatePhase `json:"phases"`
}

type ProjecttemplatePhase struct {
	ID          string                 `json:"id"`
	PhaseOrder  byte                   `json:"phase_order"`
	Name        string                 `json:"name" csval:"req"`
	Description string                 `json:"description" csval:"req"`
	Tasks       []*ProjectTemplateTask `json:"tasks"`
}

type ProjectTemplateTask struct {
	ID              string `json:"id"`
	Name            string `json:"name" csval:"req"`
	Description     string `json:"description" csval:"req"`
	RequiredSkillID string `json:"required_skill_id" csval:"req"`
}
