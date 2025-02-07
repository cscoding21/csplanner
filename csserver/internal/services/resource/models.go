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
	Name                  string                        `json:"name" caval:"req"`
	Type                  resourcetype.ResourceType     `json:"type" caval:"req"`
	Status                resourcestatus.ResourceStatus `json:"status" caval:"req"`
	RoleID                *string                       `json:"role_id"`
	UserEmail             *string                       `json:"user_email"`
	ProfileImage          *string                       `json:"profile_image"`
	InitialCost           float64                       `json:"initial_cost"`
	AnnualizedCost        float64                       `json:"annualized_cost"`
	Skills                []*Skill                      `json:"skills"`
	AvailableHoursPerWeek int                           `json:"available_hours_per_week"`
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

// GetHourlyRate return the hourly charge for a user
func (r *Resource) GetHourlyRate(roleMap map[string]Role, defaultRate float64, hoursPerYear int) float64 {
	calculated := 0.0
	if r.AnnualizedCost > 0 && hoursPerYear > 0 {
		calculated = r.AnnualizedCost / float64(hoursPerYear)
	}

	if calculated > 0.0 {
		return calculated
	}

	role, ok := roleMap[*r.RoleID]
	if ok && *role.HourlyRate > 0.0 {
		return *role.HourlyRate
	}

	return defaultRate
}

// Skill for people resources...a skill and proficiency that they possess
type Skill struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Proficiency *float64 `json:"proficiency"`
}

type Role struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	Name        string   `json:"name"`
	Description string   `json:"description"`
	HourlyRate  *float64 `json:"hourly_rate"`
}
