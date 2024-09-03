package interfaces

import "context"

type PubSubClient interface {
	Publish(ctx context.Context, service string, object string, eventName string, body interface{}) error
}
