package tests

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
)

// getTestContext return a context object suitable to running test against live services
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