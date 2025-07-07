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
	Body          any
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
func (s *PubSubProvider) GetSubjectName(ctx context.Context, service string, object string, eventName string) string {
	orgKey := config.GetOrgUrlKeyFromContext(ctx)
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

	return client //pubSubClient
}

// GetPubSubConnDurable return a configured Jetstream connection
func (s *PubSubProvider) getStreamContext(ctx context.Context) jetstream.JetStream {
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

// GetStreamConsumer
// func (s *PubSubProvider) GetStreamConsumer(ctx context.Context, consumerName string) (jetstream.Consumer, error) {
// 	js := *s.GetPubSubConnDurable(ctx)

// 	cons, err := js.(ctx, STREAM_NAME, jetstream.ConsumerConfig{
// 		Durable:   consumerName,
// 		AckPolicy: jetstream.AckExplicitPolicy,
// 	})
// 	if err != nil {
// 		log.Errorf("Error creating consumer: %s", err)
// 		return nil, err
// 	}

// 	log.Infof("Created consumer %s on stream %s", consumerName, STREAM_NAME)

// 	return cons, nil
// }

// Publish publish a message to the pubsub platform
func (s *PubSubProvider) Publish(ctx context.Context, service string, object string, eventName string, body any) error {
	client := s.GetPubSubConn()
	//defer client.Drain()

	subject := s.GetSubjectName(ctx, service, object, eventName)
	data, err := s.getWrappedData(body)
	if err != nil {
		log.Error(err)
		return err
	}

	client.Publish(subject, data)

	return nil
}

// StreamPublish publisn an event to a durable stream
func (s *PubSubProvider) StreamPublish(ctx context.Context, service string, object string, eventName string, body any) error {
	js := s.getStreamContext(ctx)
	streamName := strings.ToLower(config.Config.PubSub.StreamName)
	orgName := strings.ToLower(config.GetOrgUrlKeyFromContext(ctx))

	subject := fmt.Sprintf("%s.%s.%s.%s.%s", streamName, orgName, service, object, eventName)

	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     streamName,
		Subjects: []string{fmt.Sprintf("%s.*.*.*.>", streamName)},
	})
	if err != nil {
		return err
	}

	_, err = js.Publish(ctx, subject, jsonData)

	log.Info(subject)
	log.Info(jsonData)

	return err
}

// Subscribe consumes messages from NATS
func (s *PubSubProvider) Subscribe(
	ctx context.Context,
	service string,
	object string,
	eventName string) (*nats.Subscription, error) {
	conn := s.GetPubSubConn()
	subject := s.GetSubjectName(ctx, service, object, eventName)

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
