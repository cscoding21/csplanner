package rtypes

type ResourceStatus string
type ResourceType string

const (
	ResourceStatusNotSet ResourceStatus = "notset"

	Inhouse  = "inhouse"
	Proposed = "proposed"
)

const (
	ResourceTypeNotSet ResourceType = "notset"

	Human     = "human"
	Software  = "software"
	Equipment = "equipment"
)
