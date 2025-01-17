//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package schedule

import (
	"csserver/internal/common"
	"time"

	"github.com/google/uuid"
)

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
	Type    ScheduleExceptionType
	Level   ScheduleExceptionLevel
	Scope   string
	Message string
}

type CalculatedScheduleInfo struct {
	IsComplete bool
}
