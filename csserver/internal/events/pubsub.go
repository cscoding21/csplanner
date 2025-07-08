package events

import (
	"context"
	"csserver/internal/config"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	log "github.com/sirupsen/logrus"
)

// var pubSubClient *nats.Conn
var pubSubClientDurable jetstream.JetStream

// CSPLANNER.[service].[object].[eventname]
const SUBJECT_FORMAT = "%s.%s.%s.%s"

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
	OrgKey    string
	ID        string
	Timestamp time.Time
	Body      any
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
func (s *PubSubProvider) GetSubjectName(orgKey string, service string, object string, eventName string) string {
	return fmt.Sprintf(s.SubjectFormat, s.StreamName, orgKey, service, object, eventName)
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

	return client
}

// GetPubSubConnDurable return a configured Jetstream connection
func (s *PubSubProvider) getStreamContext() jetstream.JetStream {
	if pubSubClientDurable != nil {
		return pubSubClientDurable
	}

	conn := s.GetPubSubConn()
	if conn == nil {
		return nil
	}

	js, err := jetstream.New(conn)
	if err != nil {
		log.Fatal(err)
	}

	pubSubClientDurable = js

	return pubSubClientDurable
}

// Publish publish a message to the pubsub platform
func (s *PubSubProvider) Publish(ctx context.Context, orgName string, service string, object string, eventName string, body any) error {
	client := s.GetPubSubConn()
	//defer client.Drain()

	subject := s.GetSubjectName(orgName, service, object, eventName)
	data, err := s.getWrappedData(orgName, body)
	if err != nil {
		log.Error(err)
		return err
	}

	client.Publish(subject, data)

	return nil
}

// StreamPublish publisn an event to a durable stream
func (s *PubSubProvider) StreamPublish(ctx context.Context, orgName string, service string, object string, eventName string, body any) error {
	js := s.getStreamContext()
	streamName := strings.ToLower(config.Config.PubSub.StreamName)

	subject := fmt.Sprintf("%s.%s.%s.%s.%s", streamName, orgName, service, object, eventName)

	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	wrappedData, err := s.getWrappedData(orgName, body)
	if err != nil {
		return err
	}

	//sj := fmt.Sprintf("%s.*.*.*.>", orgName)
	_, err = js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     streamName,
		Subjects: []string{subject},
	})
	if err != nil {
		return err
	}

	_, err = js.Publish(ctx, subject, wrappedData)

	log.Info(subject)
	log.Info(jsonData)

	return err
}

// StreamConsume read messages from a durable stream
func (s *PubSubProvider) GetStreamConsumer(ctx context.Context, name string, subject string) (jetstream.Consumer, error) {
	js := s.getStreamContext()

	cons, err := js.CreateOrUpdateConsumer(ctx, strings.ToLower(s.StreamName), jetstream.ConsumerConfig{
		Durable: name,
		// FilterSubject: subject,
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Errorf("Error creating consumer: %s", err)
		return nil, err
	}

	log.Infof("Created consumer %s on stream %s", subject, s.StreamName)

	return cons, nil
}

// Subscribe consumes messages from NATS
func (s *PubSubProvider) Subscribe(
	ctx context.Context,
	orgName string,
	service string,
	object string,
	eventName string) (*nats.Subscription, error) {
	conn := s.GetPubSubConn()
	subject := s.GetSubjectName(orgName, service, object, eventName)

	sub, err := conn.Subscribe(subject, func(msg *nats.Msg) {
		log.Debugf("Received message on subject %s: %s", subject, string(msg.Data))
	})
	if err != nil {
		log.Errorf("Subscribe Error: %s", err)
		return nil, err
	}

	return sub, nil
}

func (s *PubSubProvider) getWrappedData(org string, data any) ([]byte, error) {
	msg := MessageWrapper{
		OrgKey:    org,
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
