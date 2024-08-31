package notification

type NotificationType int

const (
	NotSet NotificationType = iota
	Mention
	Reply
)
