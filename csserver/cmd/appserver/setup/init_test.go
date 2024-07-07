package setup

import (
	"context"
	"csserver/internal/config"
	"testing"
)

func init() {
	config.InitConfig()
	config.InitLogger()
}

func TestSetupTestData(t *testing.T) {
	ctx := context.Background()
	err := SetupTestData(ctx)
	if err != nil {
		t.Error(err)
	}
}
