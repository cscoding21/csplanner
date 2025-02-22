package schedule

// type ScheduleExceptionType string
// type ScheduleExceptionLevel byte

const (
	ScheduleCannotComplete           = ScheduleExceptionType("ScheduleCannotComplete")
	TaskLacksAssignedResource        = ScheduleExceptionType("TaskLacksAssignedResource")
	TaskAssignedResourceMissingSkill = ScheduleExceptionType("TaskAssignedResourceMissingSkill")
	AssignedResourceNotOnStaff       = ScheduleExceptionType("AssignedResourceNotOnStaff")
)

const (
	ScheduleError   = ScheduleExceptionLevel(1)
	ScheduleWarning = ScheduleExceptionLevel(2)
	ScheduleInfo    = ScheduleExceptionLevel(3)
)
