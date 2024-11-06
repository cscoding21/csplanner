package tests

import (
	"context"
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

func getTestContext() context.Context {
	service := factory.GetUserService()
	ctx := context.Background()

	anonID, err := service.GetUser(config.Config.Default.BotUserEmail)
	if err != nil {
		panic(err)
	}

	ctx = context.WithValue(ctx, config.UserEmailKey, config.Config.Default.BotUserEmail)
	ctx = context.WithValue(ctx, config.UserIDKey, anonID)

	return ctx
}
