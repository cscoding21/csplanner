//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package project

import (
	"csserver/internal/common"
	"time"
)

type Project struct {
	//---common for all DB objects
	common.ControlFields `json:"control_fields" csval:"validate"`

	//---TODO: add fields here
	ProjectBasics     *ProjectBasics      `json:"basics"`
	ProjectValue      *ProjectValue       `json:"value"`
	ProjectCost       *ProjectCost        `json:"cost"`
	ProjectDaci       *ProjectDaci        `json:"daci"`
	ProjectFeatures   []*ProjectFeature   `json:"features"`
	ProjectMilestones []*ProjectMilestone `json:"milestones"`
}

type ProjectBasics struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      ProjectState `json:"status"`
	StartDate   *time.Time   `json:"start_time"`
	OwnerID     string       `json:"owner_id"`
}

type ProjectValue struct {
	FundingSource        string  `json:"funding_source"`
	DiscountRate         float64 `json:"discount_rate"`
	YearOneValue         float64 `json:"year_one_value"`
	YearTwoValue         float64 `json:"year_two_value"`
	YearThreeValue       float64 `json:"year_three_value"`
	YearFourValue        float64 `json:"year_four_value"`
	YearFiveValue        float64 `json:"year_five_value"`
	NetPresentValue      float64 `json:"net_present_value"`
	InternalRateOfReturn float64 `json:"internal_rate_of_return"`
}

type ProjectCost struct {
	Ongoing      *float64 `json:"ongoing"`
	BlendedRate  *float64 `json:"blended_rate"`
	InitialCost  float64  `json:"initial_cost"`
	HourEstimate int      `json:"hours_estimate"`
}

type ProjectDaci struct {
	Driver      []*string `json:"driver"`
	Approver    []*string `json:"approver"`
	Contributor []*string `json:"contributor"`
	Informed    []*string `json:"informed"`
}

type ProjectFeature struct {
	ID          *string         `json:"id,omitempty"`
	Priority    FeaturePriority `json:"priority"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Status      string          `json:"status"`
}

type ProjectMilestone struct {
	ID        *string                 `json:"id,omitempty"`
	StartDate *time.Time              `json:"start_date"`
	EndDate   *time.Time              `json:"end_date"`
	Phase     *ProjectMilestonePhase  `json:"phase"`
	Tasks     []*ProjectMilestoneTask `json:"tasks"`
}

type ProjectMilestonePhase struct {
	ID          string `json:"id,omitempty"`
	Order       byte   `json:"order"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectMilestoneTask struct {
	ID               *string         `json:"id,omitempty"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	HourEstimate     int             `json:"hour_estimate"`
	ResourceIDs      []string        `json:"resources_ids"`
	RequiredSkillIDs []string        `json:"required_skill_ids"`
	Status           MilestoneStatus `json:"status"`
	StartDate        *time.Time      `json:"start_date"`
	EndDate          *time.Time      `json:"end_date"`
}