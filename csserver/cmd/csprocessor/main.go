package main

import (
	"context"
	"csserver/internal/appserv/csmap"
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"csserver/internal/events"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	log "github.com/sirupsen/logrus"
)

// setup global configuration values
func init() {
	config.InitConfig()
	config.InitLogger()

	csmap.ClearAllCaches()
}

func main() {
	log.Info("Processor started")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ps, err := factory.GetPubSubClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go processActivities(ctx, ps)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func processActivities(ctx context.Context, ps events.PubSubProvider) (jetstream.ConsumeContext, error) {
	c, err := ps.GetStreamConsumer(ctx, "activity", "csplanner.*")
	if err != nil {
		log.Fatalf("GetStreamConsumer error: %v", err)
	}

	handler, err := c.Consume(func(msg jetstream.Msg) {
		actx, err := getEventContext(msg)
		if err != nil {
			log.Errorf("getEventcontext error: %s", err)
			return
		}

		fmt.Println("----------------- MESSAGE -----------------")
		fmt.Println(msg.Subject())
		fmt.Println(string(msg.Data()))

		activityService := factory.GetActivityService(actx)
		ps := factory.GetProjectService(actx)
		rs := factory.GetResourceService(actx)

		ierr := activityService.LogActivity(actx, *ps, *rs, msg.Subject(), msg.Data())

		if ierr != nil {
			log.Errorf("LogActivity error: %s", ierr)
		}

		msg.Ack()
	})
	if err != nil {
		fmt.Println(err)
	}

	return handler, err
}

func getEventContext(msg jetstream.Msg) (context.Context, error) {
	ctx := context.Background()

	var env events.MessageWrapper
	err := json.Unmarshal(msg.Data(), &env)
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, config.OrgUrlKey, env.OrgKey)
	ctx = context.WithValue(ctx, config.UserEmailKey, env.UserEmail)
	ctx = context.WithValue(ctx, config.UserIDKey, env.UserID)

	return ctx, nil
}
