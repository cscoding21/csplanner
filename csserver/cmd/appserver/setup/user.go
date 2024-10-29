package setup

import (
	"context"

	"csserver/internal/appserv/factory"
	"csserver/internal/config"

	userPackage "csserver/internal/services/iam/user"

	log "github.com/sirupsen/logrus"
)

func CreateOrGetBot(ctx context.Context) *userPackage.User {
	service := factory.GetIAMAdminService()

	user, err := service.GetUser(ctx, config.Config.Default.BotUserEmail)
	if err != nil {
		log.Error(err)
	}

	if user != nil && len(user.Email) > 0 {
		return user
	}

	botUser := userPackage.User{
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

	usrs := []userPackage.User{
		//---users
		{Email: "jeph@jmk21.com", FirstName: "Geomfry", LastName: "McHale", Password: "localpass", ConfirmPassword: "localpass"},
		{Email: "tifa@jmk21.com", FirstName: "Tifa", LastName: "Lockhart", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/tifa.png"},
		{Email: "cloud@jmk21.com", FirstName: "Cloud", LastName: "Strife", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/cloud.png"},
		{Email: "aerith@jmk21.com", FirstName: "Aerith", LastName: "Gainsborough", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/aerith.png"},
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
