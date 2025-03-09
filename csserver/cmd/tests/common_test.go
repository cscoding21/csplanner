package tests

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
)

// getTestContext return a context object suitable to running test against live services
func getTestContext() context.Context {
	service := factory.GetAppuserService()
	ctx := context.Background()
	anonID := ""

	anonUser, err := service.GetAppuser(ctx, config.Config.Default.BotUserEmail)
	if err == nil {
		anonID = anonUser.ID
	}

	ctx = context.WithValue(ctx, config.UserEmailKey, config.Config.Default.BotUserEmail)
	ctx = context.WithValue(ctx, config.UserIDKey, anonID)

	return ctx
}
