package setup

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/appserv/orgmap"
	"csserver/internal/config"
	"errors"

	log "github.com/sirupsen/logrus"
)

const TEST_KEY = "localhost"

func SetupTestData(ctx context.Context) error {
	ctx = context.WithValue(ctx, config.OrgUrlKey, TEST_KEY)

	ctx = GetBotContext(ctx)
	org, err := orgmap.GetSaaSOrg(ctx)
	if err != nil {
		log.Fatalf("ERROR setting up test org: %s\n", err)
	} else {
		log.Infof("Setting up data for test org %s", org.Info.Org.Name)
	}

	err = errors.Join(nil, CreateTestUsers(ctx))
	err = errors.Join(err, CreateDefaultOrganization(ctx))
	err = errors.Join(err, CreateTestRoles(ctx))
	err = errors.Join(err, CreateTestLists(ctx))
	err = errors.Join(err, CreateTestTemplates(ctx))
	err = errors.Join(err, CreateTestResources(ctx))
	err = errors.Join(err, CreateTestProjects(ctx))
	err = errors.Join(err, CreateTestComments(ctx))

	processor := factory.GetProcessorService(ctx)

	err = errors.Join(err, processor.ProcessNightly(ctx))

	return err
}

func GetBotContext(ctx context.Context) context.Context {
	user := CreateOrGetBot(ctx)

	ctx = context.WithValue(ctx, config.UserEmailKey, config.Config.Default.BotUserEmail)
	ctx = context.WithValue(ctx, config.UserIDKey, user.ID)

	return ctx
}
