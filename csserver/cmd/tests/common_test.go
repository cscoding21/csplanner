package tests

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/appserv/orgmap"
	"csserver/internal/config"

	log "github.com/sirupsen/logrus"
)

// getTestContext return a context object suitable to running test against live services
func getTestContext() context.Context {
	ctx := context.Background()

	ctx = context.WithValue(ctx, config.OrgUrlKey, "localhost")

	orgInfo, err := orgmap.GetSaaSOrg(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Warnf("Test org loaded: %s\n", orgInfo.Info.Org.Name)
	}

	service := factory.GetAppuserService(ctx)

	anonID := ""

	anonUser, err := service.GetAppuser(ctx, config.Config.Default.BotUserEmail)
	if err == nil {
		anonID = anonUser.ID
	}

	ctx = context.WithValue(ctx, config.UserEmailKey, config.Config.Default.BotUserEmail)
	ctx = context.WithValue(ctx, config.UserIDKey, anonID)

	return ctx
}
