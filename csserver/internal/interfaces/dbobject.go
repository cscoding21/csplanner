package interfaces

type DBObject interface {
	SetCreateInfo(string)
	SetUpdateInfo(string)
	SetDeleteInfo(string)
	HasID() bool
	GetID() string
}
