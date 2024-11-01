package resourcetype

type ResourceType string

const (
	ResourceTypeNotSet ResourceType = "notset"

	Human     = "human"
	Software  = "software"
	Equipment = "equipment"
)
