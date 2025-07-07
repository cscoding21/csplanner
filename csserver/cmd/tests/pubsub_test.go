package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"fmt"
	"testing"
)

func init() {
	config.InitConfig()
}

func TestGetConsumer(t *testing.T) {
	ctx := getTestContext()

	ps, err := factory.GetPubSubClient(ctx)
	if err != nil {
		t.Errorf("Get client error: %v", err)
	}

	to := "hello"
	err = ps.Publish(ctx, "test", "test", "publish", to)
	if err != nil {
		t.Errorf("Publish error: %v", err)
	}

	err = ps.StreamPublish(ctx, "test", "stream", "publish", to)
	if err != nil {
		t.Errorf("Stream publish error: %v", err)
	}

	fmt.Printf("HOST: %s\n", ps.Name)
}
