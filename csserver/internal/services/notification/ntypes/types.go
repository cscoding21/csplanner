package ntypes

type NotificationType int

const (
	NotSet NotificationType = iota
	Mention
	Reply
)
