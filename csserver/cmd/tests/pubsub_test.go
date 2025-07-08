package tests

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/nats-io/nats.go/jetstream"
)

func init() {
	config.InitConfig()
}

type CSBody struct {
	Content string `json:"content"`
}

const OrgName = "localhost"

func TestGetConsumer(t *testing.T) {
	//ctx := getTestContext()
	ctx := context.Background()

	ps, err := factory.GetPubSubClient(ctx)
	if err != nil {
		t.Errorf("Get client error: %v", err)
	}

	to := CSBody{Content: fmt.Sprintf("hello %d", 999)}
	err = ps.Publish(ctx, OrgName, "test", "test", "publish", to)
	if err != nil {
		t.Errorf("Publish error: %v", err)
	}

	for i := range 10 {
		to = CSBody{Content: fmt.Sprintf("hello %d", i)}
		err = ps.StreamPublish(ctx, OrgName, "test", "stream", "publish", to)
		if err != nil {
			t.Errorf("Stream publish error: %v", err)
		}
	}

	fmt.Printf("HOST: %s\n", ps.Name)

	fmt.Println("--------------------------------------")

	c, err := ps.GetStreamConsumer(ctx, "CONS", "csplanner.*")
	if err != nil {
		t.Errorf("GetStreamConsumer error: %v", err)
	}

	go handleStream(c)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println(c)
}

func handleStream(js jetstream.Consumer) error {
	_, err := js.Consume(func(msg jetstream.Msg) {
		fmt.Println("----------------- MESSAGE -----------------")
		fmt.Println(msg.Subject())
		fmt.Println(string(msg.Data()))
		msg.Ack()
	})
	if err != nil {
		fmt.Println(err)
	}
	//defer handle.Stop()

	return err
}
