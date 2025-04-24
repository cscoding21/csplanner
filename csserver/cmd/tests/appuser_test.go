package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/appserv/orgmap"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/services/iam/appuser"
	"fmt"
	"testing"
)

func TestCreateOrGetBot(t *testing.T) {
	ctx := getTestContext()

	service := factory.GetIAMAdminService(ctx)
	orgInfo, err := orgmap.GetSaaSOrg(ctx)
	if err != nil {
		t.Error(err)
	}

	user, err := service.GetUser(ctx, orgInfo.Info.Org.Realm, config.Config.Default.BotUserEmail)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(user)

	botUser := appuser.Appuser{
		ControlFields: common.ControlFields{
			ID: "user:bot",
		},
		Email:           config.Config.Default.BotUserEmail,
		FirstName:       "AI",
		LastName:        "Bot",
		Password:        config.Config.Default.BotPassword,
		ConfirmPassword: config.Config.Default.BotPassword,
		ProfileImage:    "/aibot.jpg",
	}

	ur, err := service.CreateUser(ctx, orgInfo.Info.Org.Realm, &botUser)
	if err != nil {
		t.Errorf("error creating user: %v", err)
	}

	if !ur.ValidationResult.Pass {
		t.Error(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
	}
}
