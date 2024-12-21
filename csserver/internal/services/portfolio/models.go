package portfolio

import (
	"csserver/internal/calendar"
	"csserver/internal/services/project/ptypes/projectstatus"
	"time"

	"github.com/google/uuid"
)

type ResourceUtilizationTable struct {
	Resources []ResourceUtilizationItem `json:"resource_snapshot"`
}

type ResourceUtilizationItem struct {
	ProjectID          string                     `json:"project_id"`
	ProjectName        string                     `json:"project_name"`
	ProjectStatus      projectstatus.ProjectState `json:"project_status"`
	MilestoneName      string                     `json:"milestone_name"`
	MilestoneStartDate *time.Time                 `json:"milestone_start_date"`
	MilestoneEndDate   *time.Time                 `json:"milestone_end_date"`
	TaskName           string                     `json:"task_name"`
	TaskStartDate      *time.Time                 `json:"task_start_date"`
	TaskEndDate        *time.Time                 `json:"task_end_date"`
	TaskHourEstimate   int                        `json:"task_hour_estimate"`
	ResourceID         string                     `json:"resource_id"`
	ResourceName       string                     `json:"resource_name"`
	AllocationPercent  float64                    `json:"allocation_percent"`
}

type ProjectSchedule struct {
	ProjectID            string
	ProjectName          string
	ProjectActivityWeeks *[]ProjectActivityWeek
	Exceptions           []string
}

type ProjectActivityWeek struct {
	Week       calendar.CSWeek
	Activities []ProjectActivity
}

type ProjectActivity struct {
	ProjectID     string
	ProjectName   string
	MilestoneID   string
	MilestoneName string
	TaskID        string
	TaskName      string
	ResourceID    string
	ResourceName  string
	HoursSpent    int
}

type ScheduleBatch struct {
	Order           int
	BatchID         uuid.UUID
	HoursToComplete int
	ScheduleItems   []ScheduleWorkspace
}

type ScheduleWorkspace struct {
	MilestoneID     string
	MilestoneName   string
	TaskID          string
	TaskName        string
	ResourceIDs     []string
	HoursToSchedule int
	HoursScheduled  int
}
