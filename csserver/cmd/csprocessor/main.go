package main

import (
	"context"
	"csserver/internal/appserv/csmap"
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
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

	log.Info(ps.StreamName)
	//sub, err := ps.Subscribe(ctx, )

	forever()
	select {}
}

func example() {
	// connect to nats server
	nc, _ := nats.Connect(config.Config.PubSub.Host)

	// create jetstream context from nats connection
	js, _ := jetstream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// get existing stream handle
	stream, _ := js.Stream(ctx, "foo")

	// retrieve consumer handle from a stream
	cons, _ := stream.Consumer(ctx, "cons")

	// consume messages from the consumer in callback
	cc, _ := cons.Consume(func(msg jetstream.Msg) {
		fmt.Println("Received jetstream message: ", string(msg.Data()))
		msg.Ack()
	})
	defer cc.Stop()
}

func forever() {
	for {
		fmt.Printf("FOREVER: %v+\n", time.Now())
		time.Sleep(time.Second * 10)
	}
}
