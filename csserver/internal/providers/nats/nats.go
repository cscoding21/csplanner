package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"csserver/internal/config"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	//"github.com/nats-io/nats.go/jetstream"
)

// var pubSubClient *nats.Conn
// var pubSubClientDurable *nats.JetStreamContext

// CSPLANNER.[service].[object].[eventname]
const SUBJECT_FORMAT = "%s.%s.%s.%s"
const STREAM_NAME = "CSPLANNER"

type MessageWrapper struct {
	ID        string
	Timestamp time.Time
	Body      interface{}
}

// getSubjectName enforce a standard around pubsub subject naming
func GetSubjectName(service string, object string, eventName string) string {
	return fmt.Sprintf(SUBJECT_FORMAT, STREAM_NAME, service, object, eventName)
}

// GetPubSubClient return a configured NATS client
func GetPubSubClient(ctx context.Context) *nats.Conn {
	client, err := nats.Connect(
		config.Config.PubSub.Host,
		nats.Name("analyzer-nats"),
		nats.Timeout(10*time.Second))
	if err != nil {
		log.Errorf("NATS Connect Error: %s", err)
		return nil
	}

	return client //pubSubClient
}

// Publish publish a message to the pubsub platform
func Publish(ctx context.Context, service string, object string, eventName string, body interface{}) error {
	client := GetPubSubClient(ctx)
	//defer client.Drain()

	subject := GetSubjectName(service, object, eventName)
	data, err := getWrappedData(body)
	if err != nil {
		log.Error(err)
		return err
	}

	client.Publish(subject, data)

	return nil
}

// Subscribe consumes messages from NATS
func Subscribe(ctx context.Context, service string, object string, eventName string) (*nats.Subscription, error) {
	client := GetPubSubClient(ctx)
	subject := GetSubjectName(service, object, eventName)

	sub, err := client.Subscribe(subject, func(msg *nats.Msg) {
		log.Debugf("Received message on subject %s: %s", subject, string(msg.Data))
	})
	if err != nil {
		log.Errorf("Subscribe Error: %s", err)
		return nil, err
	}

	return sub, nil
}

func getWrappedData(data interface{}) ([]byte, error) {
	msg := MessageWrapper{
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Body:      data,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		log.Errorf("Error marshalling data: %s", err)
		return nil, err
	}

	return b, nil
}
