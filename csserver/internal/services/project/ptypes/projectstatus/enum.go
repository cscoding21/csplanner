package projectstatus

type ProjectState string

const (
	ProjectStateNotSet ProjectState = "notset"

	Exception  ProjectState = "exception"  //---denotes a mapping error...should never happen
	NewProject ProjectState = "new"        //---when project first created
	Draft      ProjectState = "draft"      //---project details (features, cost, value) are being fleshed out
	Proposed   ProjectState = "proposed"   //---project details are complete and is ready for review
	Approved   ProjectState = "approved"   //---project has been reviewed and approved
	Rejected   ProjectState = "rejected"   //---project has been reviewed and rejected
	Backlogged ProjectState = "backlogged" //---project is approved and awaiting resource availability
	Scheduled  ProjectState = "scheduled"  //---project has been staffed and a start date has been set
	InFlight   ProjectState = "inflight"   //---project has been started
	Complete   ProjectState = "complete"   //---project has been completed
	Deferred   ProjectState = "deferred"   //---project was in flight, but work has been postponed
	Abandoned  ProjectState = "abandoned"  //---project made it past scheduled, but has been cancelled
)
