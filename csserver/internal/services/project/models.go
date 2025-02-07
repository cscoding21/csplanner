//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package project

import (
	"csserver/internal/common"
	"csserver/internal/services/project/ptypes/featurepriority"
	"csserver/internal/services/project/ptypes/milestonestatus"
	"csserver/internal/services/project/ptypes/projectstatus"
	"time"

	"github.com/cscoding21/csval/validate"
)

// Project the primary object representing a product
type Project struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	ProjectBasics      *ProjectBasics        `json:"basics"`
	ProjectStatusBlock *ProjectStatusBlock   `json:"status"`
	ProjectValue       *ProjectValue         `json:"value"`
	ProjectCost        *ProjectCost          `json:"cost"`
	ProjectDaci        *ProjectDaci          `json:"daci"`
	ProjectFeatures    []*ProjectFeature     `json:"features"`
	ProjectMilestones  []*ProjectMilestone   `json:"milestones"`
	Calculated         ProjectCalculatedData `json:"calculated"`
}

// ProjectBasics basic elements of a project
type ProjectBasics struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	StartDate   *time.Time `json:"start_time"`
	OwnerID     string     `json:"owner_id"`
}

// ProjectValue properties used in calculating the value of a project
type ProjectValue struct {
	DiscountRate      float64                    `json:"discount_rate"`
	IsCapitalized     bool                       `json:"is_capitalized"`
	ProjectValueLines []*ProjectValueLine        `json:"project_value_lines"`
	Calculated        ProjectValueCalculatedData `json:"calculated"`
}

type ProjectValueLine struct {
	ID             *string `json:"id,omitempty"`
	ValueCategory  string  `json:"value_category"`
	FundingSource  string  `json:"funding_source"`
	YearOneValue   float64 `json:"year_one_value"`
	YearTwoValue   float64 `json:"year_two_value"`
	YearThreeValue float64 `json:"year_three_value"`
	YearFourValue  float64 `json:"year_four_value"`
	YearFiveValue  float64 `json:"year_five_value"`
}

// ProjectCost properties used to calculate the cost of a project
type ProjectCost struct {
	Ongoing     *float64                  `json:"ongoing"`
	BlendedRate *float64                  `json:"blended_rate"`
	Calculated  ProjectCostCalculatedData `json:"calculated"`
}

// ProjectDaci members of a project team represented using the DACI model
type ProjectDaci struct {
	DriverIDs      []*string `json:"driver_ids"`
	ApproverIDs    []*string `json:"approver_ids"`
	ContributorIDs []*string `json:"contributor_ids"`
	InformedIDs    []*string `json:"informed_ids"`
}

type ProjectFeature struct {
	ID          *string                         `json:"id,omitempty"`
	Priority    featurepriority.FeaturePriority `json:"priority"`
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	Status      string                          `json:"status"`
}

type ProjectMilestone struct {
	ID         *string                 `json:"id,omitempty"`
	Phase      *ProjectMilestonePhase  `json:"phase"`
	Tasks      []*ProjectMilestoneTask `json:"tasks"`
	Calculated ProjectMilestoneCalculatedData
}

type ProjectMilestonePhase struct {
	ID          string `json:"id,omitempty"`
	Order       byte   `json:"order"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectMilestoneTask struct {
	ID              *string                         `json:"id,omitempty"`
	Name            string                          `json:"name"`
	Description     string                          `json:"description"`
	HourEstimate    int                             `json:"hour_estimate"`
	ResourceIDs     []string                        `json:"resources_ids"`
	RequiredSkillID string                          `json:"required_skill_id"`
	Status          milestonestatus.MilestoneStatus `json:"status"`
	Calculated      ProjectTaskCalculatedData       `json:"calculated"`
}

type ProjectCalculatedData struct {
	Team []string `json:"team"`
}

type ProjectCostCalculatedData struct {
	InitialCost     float64 `json:"initial_cost"`
	HourEstimate    int     `json:"hours_estimate"`
	HoursActualized int     `json:"hours_actualized"`
}

type ProjectValueCalculatedData struct {
	NetPresentValue      float64 `json:"net_present_value"`
	InternalRateOfReturn float64 `json:"internal_rate_of_return"`
	YearOneValue         float64 `json:"year_one_value"`
	YearTwoValue         float64 `json:"year_two_value"`
	YearThreeValue       float64 `json:"year_three_value"`
	YearFourValue        float64 `json:"year_four_value"`
	YearFiveValue        float64 `json:"year_five_value"`
}

type ProjectMilestoneCalculatedData struct {
	TotalHours     int  `json:"total_hours"`
	HoursRemaining int  `json:"hours_remaining"`
	IsComplete     bool `json:"is_complete"`
	IsInFlight     bool `json:"is_in_flight"`
	TotalTasks     int  `json:"total_tasks"`
	CompletedTasks int  `json:"completed_tasks"`
	RemovedHours   int  `json:"removed_hours"`
}

type ProjectTaskCalculatedData struct {
	ParallelWorkThreads        float64 `json:"parallel_resources"`
	ActualizedHoursToComplete  int     `json:"actualized_hours_to_complete"`
	DaysToComplete             int     `json:"days_to_complete"`
	CommunicationOverhead      int     `json:"communication_overhead"`
	ResourceSaturationOverhead int     `json:"resource_saturation_overhead"`
	ActualizedCost             float64 `json:"actualized_cost"`
	AverageHourlyRate          float64 `json:"average_hourly_rate"`
	ResourceContention         float64 `json:"resource_contention"`
	SkillsHourAdjustment       int     `json:"skills_hour_adjustment"`
	CommsHourAdjustment        int     `json:"comms_hour_adjustment"`

	Exceptions []string `json:"exceptions"`
}

type ProjectStatusBlock struct {
	Status            projectstatus.ProjectState `json:"status"`
	AllowedNextStates []*ProjectStatusTransition
}

type ProjectStatusTransition struct {
	NextState    projectstatus.ProjectState `json:"next_state"`
	CanEnter     bool                       `json:"can_enter"`
	CheckResults validate.ValidationResult  `json:"check_results"`
}
