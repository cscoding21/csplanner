package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"testing"
)

func init() {
	config.InitConfig()
}

func TestPortfolioGetResourceUtilizationTable(t *testing.T) {
	service := factory.GetPortfolioService()
	ctx := getTestContext()

	table, err := service.GetResourceUtilizationTable(ctx)
	if err != nil {
		t.Error(err)
	}

	t.Log(table)
}
