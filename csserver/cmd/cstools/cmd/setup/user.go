package setup

import (
	"context"

	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/config"

	"csserver/internal/services/iam/appuser"

	log "github.com/sirupsen/logrus"
)

func CreateOrGetBot(ctx context.Context) *appuser.Appuser {
	service := factory.GetIAMAdminService()

	user, err := service.GetUser(ctx, config.Config.Default.BotUserEmail)
	if err != nil {
		log.Fatal(err)
	}

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

	ur, err := service.CreateUser(ctx, &botUser)
	if err != nil {
		log.Errorf("error creating user: %v", err)
	}

	if !ur.ValidationResult.Pass {
		log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
	}

	return user
}

// CreateTestUsers create a set of users for testing
func CreateTestUsers(ctx context.Context) error {
	service := factory.GetIAMAdminService()

	usrs := []appuser.Appuser{
		//---users
		{ControlFields: common.ControlFields{ID: "jeph@jmk21.com"}, Email: "jeph@jmk21.com", FirstName: "Geomfry", LastName: "McHale", Password: "localpass", ConfirmPassword: "localpass"},
		{ControlFields: common.ControlFields{ID: "tifa@jmk21.com"}, Email: "tifa@jmk21.com", FirstName: "Tifa", LastName: "Lockhart", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/tifa.png"},
		{ControlFields: common.ControlFields{ID: "cloud@jmk21.com"}, Email: "cloud@jmk21.com", FirstName: "Cloud", LastName: "Strife", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/cloud.png"},
		{ControlFields: common.ControlFields{ID: "aerith@jmk21.com"}, Email: "aerith@jmk21.com", FirstName: "Aerith", LastName: "Gainsborough", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/aerith.png"},
	}

	for _, u := range usrs {
		ur, err := service.CreateUser(ctx, &u)
		if err != nil {
			log.Errorf("NewUser(%s): %s\n", u.Email, err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	}

	return nil
}
