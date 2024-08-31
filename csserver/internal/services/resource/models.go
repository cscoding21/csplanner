//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package resource

import (
	"csserver/internal/common"
)

// Resource represents a person or purchase that contributs to the implementation of a project
type Resource struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	Name         string   `json:"name" caval:"req"`
	Type         string   `json:"type" caval:"req"`
	Role         string   `json:"role"`
	UserEmail    *string  `json:"user_email"`
	ProfileImage *string  `json:"profile_image"`
	Skills       []*Skill `json:"skills"`
	IsBot        bool     `json:"is_bot"`
}

// Skill for people resources...a skill and proficiency that they possess
type Skill struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Proficiency *float64 `json:"proficiency"`
}
