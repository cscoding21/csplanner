package portfolio

import (
	"csserver/internal/services/project/ptypes"
	"time"
)

type ResourceUtilization struct {
	Resources []ResourceUtilizationItem `json:"resource_snapshot"`
}

type ResourceUtilizationItem struct {
	ProjectID          string              `json:"project_id"`
	ProjectName        string              `json:"project_name"`
	ProjectStatus      ptypes.ProjectState `json:"project_status"`
	MilestoneName      string              `json:"milestone_name"`
	MilestoneStartDate *time.Time          `json:"milestone_start_date"`
	MilestoneEndDate   *time.Time          `json:"milestone_end_date"`
	TaskName           string              `json:"task_name"`
	TaskStartDate      *time.Time          `json:"task_start_date"`
	TaskEndDate        *time.Time          `json:"task_end_date"`
	TaskHourEstimate   int                 `json:"task_hour_estimate"`
	ResourceID         string              `json:"resource_id"`
	ResourceName       string              `json:"resource_name"`
}
