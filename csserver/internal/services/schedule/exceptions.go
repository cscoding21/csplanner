package schedule

type ScheduleExceptionType string

const (
	TaskLacksAssignedResource        = ScheduleExceptionType("TaskLacksAssignedResource")
	TaskAssignedResourceMissingSkill = ScheduleExceptionType("TaskAssignedResourceMissingSkill")
)
