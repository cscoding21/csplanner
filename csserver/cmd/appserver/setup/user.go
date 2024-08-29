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
		Email:     config.Config.Default.BotUserEmail,
		FirstName: "AI",
		LastName:  "Bot",
		Password:  "<PASSWORD>",
	}

	err = service.NewUser(ctx, &botUser)
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
		{Email: "jeph@jmk21.com", FirstName: "Geomfry", LastName: "McHale", Password: "localpass"},
		{Email: "tifa@jmk21.com", FirstName: "Tifa", LastName: "Lockhart", Password: "localpass"},
		{Email: "cloud@jmk21.com", FirstName: "Cloud", LastName: "Strife", Password: "localpass"},
		{Email: "aerith@jmk21.com", FirstName: "Aerith", LastName: "Gainsborough", Password: "localpass"},
		{Email: "barrett@jmk21.com", FirstName: "Barrett", LastName: "Wallace", Password: "localpass"},
		{Email: "jessie@jmk21.com", FirstName: "Jessie", LastName: "Raspberry", Password: "localpass"},
		{Email: "biggs@jmk21.com", FirstName: "Biggs", LastName: "Avalanche", Password: "localpass"},
		{Email: "wedge@jmk21.com", FirstName: "Wedge", LastName: "Avalanche", Password: "localpass"},
		{Email: "cid@jmk21.com", FirstName: "Cid", LastName: "Highwind", Password: "localpass"},
		{Email: "vincent@jmk21.com", FirstName: "Vincent", LastName: "Valentine", Password: "localpass"},
		{Email: "yuffie@jmk21.com", FirstName: "Yuffie", LastName: "Kisaragi", Password: "localpass"},
		{Email: "red13@jmk21.com", FirstName: "Red 13", LastName: "Nanaki", Password: "localpass"},
		{Email: "zack@jmk21.com", FirstName: "Zack", LastName: "Fair", Password: "localpass"},
	}

	for _, u := range usrs {
		err := service.NewUser(ctx, &u)
		if err != nil {
			log.Errorf("NewUser(%s): %s\n", u.Email, err)
		}
	}

	return nil
}
