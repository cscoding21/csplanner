//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package resource

import (
	"csserver/internal/common"
	"strings"

	"csserver/internal/services/resource/rtypes/resourcestatus"
	"csserver/internal/services/resource/rtypes/resourcetype"
)

// Resource represents a person or purchase that contributs to the implementation of a project
type Resource struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	Name           string                        `json:"name" caval:"req"`
	Type           resourcetype.ResourceType     `json:"type" caval:"req"`
	Status         resourcestatus.ResourceStatus `json:"status" caval:"req"`
	Role           string                        `json:"role"`
	UserEmail      *string                       `json:"user_email"`
	ProfileImage   *string                       `json:"profile_image"`
	InitialCost    float64                       `json:"initial_cost"`
	AnnualizedCost float64                       `json:"annualized_cost"`
	Skills         []*Skill                      `json:"skills"`
}

// GetSkill return a skill by its name for a given resource if it exists
func (r *Resource) GetSkill(id string) *Skill {
	for _, s := range r.Skills {
		if strings.EqualFold(s.ID, id) {
			return s
		}
	}

	return nil
}

// Skill for people resources...a skill and proficiency that they possess
type Skill struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Proficiency *float64 `json:"proficiency"`
}
