package milestonestatus

type MilestoneStatus string

const (
	MilestoneStatusNotSet MilestoneStatus = "notset"
	New                   MilestoneStatus = "new"
	Accepted              MilestoneStatus = "accepted"
	Removed               MilestoneStatus = "removed"
	Done                  MilestoneStatus = "done"
)
