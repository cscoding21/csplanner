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
		Email:    config.Config.Default.BotUserEmail,
		Name:     "AI Bot",
		Password: "<PASSWORD>",
	}

	user, err = service.CreateUser(ctx, &botUser)
	if err != nil {
		log.Errorf("error creating user: %v", err)
	}

	return user
}

// CreateTestUsers create a set of users for testing
func CreateTestUsers(ctx context.Context) error {
	service := factory.GetUserService()

	usrs := []userPackage.User{
		//---users
		{Email: "jeph@jmk21.com", Name: "Jeph", Password: "localpass"},
		{Email: "tifa@jmk21.com", Name: "Tifa", Password: "localpass"},
		{Email: "cloud@jmk21.com", Name: "Cloud", Password: "localpass"},
		{Email: "aerith@jmk21.com", Name: "Aerith", Password: "localpass"},
		{Email: "barrett@jmk21.com", Name: "Barrett", Password: "localpass"},
		{Email: "jessie@jmk21.com", Name: "Jessie", Password: "localpass"},
		{Email: "biggs@jmk21.com", Name: "Biggs", Password: "localpass"},
		{Email: "wedge@jmk21.com", Name: "Wedge", Password: "localpass"},
		{Email: "cid@jmk21.com", Name: "Cid", Password: "localpass"},
		{Email: "vincent@jmk21.com", Name: "Vincent", Password: "localpass"},
		{Email: "yuffie@jmk21.com", Name: "Yuffie", Password: "localpass"},
		{Email: "red13@jmk21.com", Name: "Red 13", Password: "localpass"},
		{Email: "zack@jmk21.com", Name: "Zack", Password: "localpass"},
	}

	for _, u := range usrs {
		//---this will fail if user already exists, which is fine
		existingUser, err := service.GetUser(ctx, u.Email)
		if err != nil {
			log.Error(err)
		}

		if existingUser != nil && len(existingUser.Email) > 0 {
			continue
		}

		service.CreateUser(ctx, &u)
	}

	return nil
}
