package events

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

// PubSubProvider
type PubSubProvider struct {
	Host          string
	Name          string
	Timeout       time.Duration
	SubjectFormat string
	StreamName    string
}

// MessageWrapper a common message wrapper abstraction
type MessageWrapper struct {
	ID            string
	Timestamp     time.Time
	Body          interface{}
	SubjectFormat string
	StreamName    string
}

// NewPubSubProvider return an configured pubsub provider
func NewPubSubProvider(host string, name string, subjectFormat string, streamName string) PubSubProvider {
	ps := PubSubProvider{
		Host:          host,
		Name:          name,
		Timeout:       10 * time.Second,
		SubjectFormat: subjectFormat,
		StreamName:    streamName,
	}

	return ps
}

// getSubjectName enforce a standard around pubsub subject naming
func (s *PubSubProvider) GetSubjectName(service string, object string, eventName string) string {
	return fmt.Sprintf(s.SubjectFormat, s.StreamName, service, object, eventName)
}

// GetPubSubClient return a configured NATS client
func (s *PubSubProvider) GetPubSubConn() *nats.Conn {
	client, err := nats.Connect(
		s.Host,
		nats.Name(s.Name),
		nats.Timeout(s.Timeout))
	if err != nil {
		log.Errorf("NATS Connect Error: %s", err)
		return nil
	}

	return client //pubSubClient
}

// Publish publish a message to the pubsub platform
func (s *PubSubProvider) Publish(ctx context.Context, service string, object string, eventName string, body interface{}) error {
	client := s.GetPubSubConn()
	//defer client.Drain()

	subject := s.GetSubjectName(service, object, eventName)
	data, err := s.getWrappedData(body)
	if err != nil {
		log.Error(err)
		return err
	}

	client.Publish(subject, data)

	return nil
}

// Subscribe consumes messages from NATS
func (s *PubSubProvider) Subscribe(
	ctx context.Context,
	service string,
	object string,
	eventName string) (*nats.Subscription, error) {
	conn := s.GetPubSubConn()
	subject := s.GetSubjectName(service, object, eventName)

	sub, err := conn.Subscribe(subject, func(msg *nats.Msg) {
		log.Debugf("Received message on subject %s: %s", subject, string(msg.Data))
	})
	if err != nil {
		log.Errorf("Subscribe Error: %s", err)
		return nil, err
	}

	return sub, nil
}

func (s *PubSubProvider) getWrappedData(data interface{}) ([]byte, error) {
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
