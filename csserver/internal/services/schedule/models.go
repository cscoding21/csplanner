//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package schedule

import (
	"csserver/internal/calendar"
	"csserver/internal/common"
	"csserver/internal/services/project/ptypes/milestonestatus"
	"time"

	"github.com/google/uuid"
)

type ScheduleExceptionType string
type ScheduleExceptionLevel byte

type Schedule struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	ProjectID            string
	ProjectName          string
	Begin                time.Time
	End                  time.Time
	ProjectActivityWeeks []*ProjectActivityWeek
	Exceptions           []ScheduleException
	Hash                 uint64                 `hash:"ignore"`
	Calculated           CalculatedScheduleInfo `hash:"ignore"`
}

type ProjectActivityWeek struct {
	Begin      time.Time
	End        time.Time
	WeekNumber int
	Year       int
	Activities []ProjectActivity
	Risks      []ScheduleException
}

type ProjectActivity struct {
	ProjectID       string
	ProjectName     string
	MilestoneID     string
	MilestoneName   string
	TaskID          string
	TaskName        string
	ResourceID      string
	ResourceName    string
	HoursSpent      int
	HourlyRate      float64
	RequiredSkillID string
	TaskEndDate     *time.Time
	Status          milestonestatus.MilestoneStatus
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
	RequiredSkillID string
	ResourceIDs     []string
	HoursToSchedule int
	HoursScheduled  int
	Status          milestonestatus.MilestoneStatus
}

type ScheduleException struct {
	Type    ScheduleExceptionType
	Level   ScheduleExceptionLevel
	Scope   string
	Message string
}

type CalculatedScheduleInfo struct {
	IsComplete bool

	PortfolioPlannedCompleteDate *time.Time
}

type ResourceProjectHourAllocation struct {
	Week                     calendar.CSWeek
	ProjectID                string
	ProjectName              string
	ResourceID               string
	ResourceName             string
	TotalResourceHours       int
	HoursAvailableForProject int
	Contention               int
}
