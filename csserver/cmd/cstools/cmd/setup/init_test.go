package setup

import (
	"csserver/internal/config"
	"testing"
)

func init() {
	config.InitConfig()
	config.InitLogger()
}

func TestSetupTestData(t *testing.T) {
	ctx := config.NewContext()
	err := SetupTestData(ctx)
	if err != nil {
		t.Error(err)
	}
}
