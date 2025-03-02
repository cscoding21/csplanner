package resourcestatus

type ResourceStatus string

const (
	ResourceStatusNotSet ResourceStatus = "notset"

	Inhouse  = "inhouse"
	Proposed = "proposed"
	Exited   = "exited"
)
