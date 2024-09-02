//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package activity

import (
	"csserver/internal/common"
)

type Activity struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	//Type     database.RelationshipType `json:"type"`
	Type     string `json:"type"`
	Summary  string `json:"summary"`
	Detail   string `json:"detail"`
	Context  string `json:"context"`
	TargetID string `json:"target_id"`
}
