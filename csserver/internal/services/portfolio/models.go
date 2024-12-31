package portfolio

import (
	"csserver/internal/services/project/ptypes/projectstatus"
	"time"

	"github.com/google/uuid"
)

type Portfolio struct {
	Schedule []ProjectSchedule
}

func (p *Portfolio) GetDateRange() (time.Time, time.Time) {
	start := time.Now().AddDate(10, 0, 0)
	end := start.AddDate(-10, 0, 0)

	for _, s := range p.Schedule {
		if s.ProjectActivityWeeks[0].Begin.Before(start) {
			start = s.ProjectActivityWeeks[0].Begin
		}

		if s.ProjectActivityWeeks[len(s.ProjectActivityWeeks)-1].End.After(end) {
			end = s.ProjectActivityWeeks[len(s.ProjectActivityWeeks)-1].End
		}
	}

	return start, end
}

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
	Begin                time.Time
	End                  time.Time
	ProjectActivityWeeks []*ProjectActivityWeek
	Exceptions           []ScheduleException
}

type ProjectActivityWeek struct {
	Begin      time.Time
	End        time.Time
	WeekNumber int
	Year       int
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

type ScheduleException struct {
	Scope   string
	Message string
}
