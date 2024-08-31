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
	Name   string                 `json:"name"`
	Phases []ProjectTemplatePhase `json:"phases"`
}

type ProjectTemplatePhase struct {
	ID          string `json:"id,omitempty"`
	PhaseOrder  byte   `json:"phase_order"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
