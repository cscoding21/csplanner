package setup

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"errors"
)

func SetupTestData(ctx context.Context) error {
	ctx = GetBotContext(ctx)

	err := errors.Join(nil, CreateTestUsers(ctx))
	err = errors.Join(err, CreateDefaultOrganization(ctx))
	err = errors.Join(err, CreateTestLists(ctx))
	err = errors.Join(err, CreateTestTemplates(ctx))
	err = errors.Join(err, CreateTestResources(ctx))
	err = errors.Join(err, CreateTestProjects(ctx))
	err = errors.Join(err, CreateTestComments(ctx))

	processor := factory.GetProcessorService()

	err = errors.Join(err, processor.ProcessNightly(ctx))

	return err
}

func GetBotContext(ctx context.Context) context.Context {
	user := CreateOrGetBot(ctx)

	ctx = context.WithValue(ctx, config.UserEmailKey, user.Email)
	ctx = context.WithValue(ctx, config.UserIDKey, user.ID)

	return ctx
}
