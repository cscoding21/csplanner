package setup

import (
	"context"

	"csserver/internal/appserv/factory"
	"csserver/internal/config"

	userPackage "csserver/internal/services/iam/user"

	log "github.com/sirupsen/logrus"
)

func CreateOrGetBot(ctx context.Context) *userPackage.User {
	service := factory.GetUserService()

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

	ur, err := service.NewUser(ctx, &botUser)
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
	service := factory.GetUserService()

	usrs := []userPackage.User{
		//---users
		{Email: "jeph@jmk21.com", FirstName: "Geomfry", LastName: "McHale", Password: "localpass", ConfirmPassword: "localpass"},
		{Email: "tifa@jmk21.com", FirstName: "Tifa", LastName: "Lockhart", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/tifa.png"},
		{Email: "cloud@jmk21.com", FirstName: "Cloud", LastName: "Strife", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/cloud.png"},
		{Email: "aerith@jmk21.com", FirstName: "Aerith", LastName: "Gainsborough", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/aerith.png"},
		{Email: "barrett@jmk21.com", FirstName: "Barrett", LastName: "Wallace", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/barrett.png"},
		{Email: "jessie@jmk21.com", FirstName: "Jessie", LastName: "Raspberry", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/jessie.png"},
		{Email: "biggs@jmk21.com", FirstName: "Biggs", LastName: "Avalanche", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/biggs.png"},
		{Email: "wedge@jmk21.com", FirstName: "Wedge", LastName: "Avalanche", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/wedge.png"},
		{Email: "cid@jmk21.com", FirstName: "Cid", LastName: "Highwind", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/cid.png"},
		{Email: "vincent@jmk21.com", FirstName: "Vincent", LastName: "Valentine", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/vincent.png"},
		{Email: "yuffie@jmk21.com", FirstName: "Yuffie", LastName: "Kisaragi", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/yuffie.png"},
		{Email: "red13@jmk21.com", FirstName: "Red 13", LastName: "Nanaki", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/red13.png"},
		{Email: "zack@jmk21.com", FirstName: "Zack", LastName: "Fair", Password: "localpass", ConfirmPassword: "localpass", ProfileImage: "/zack.png"},
	}

	for _, u := range usrs {
		ur, err := service.NewUser(ctx, &u)
		if err != nil {
			log.Errorf("NewUser(%s): %s\n", u.Email, err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	}

	return nil
}
