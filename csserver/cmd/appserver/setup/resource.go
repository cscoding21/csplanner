package setup

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"csserver/internal/utils"

	"csserver/internal/services/resource"

	log "github.com/sirupsen/logrus"
)

func CreateTestResources(ctx context.Context) error {
	us := factory.GetIAMAdminService()
	rs := factory.GetResourceService()
	exitingResources, _ := rs.FindAllResources(ctx)

	if len(exitingResources.Results) > 0 {
		return nil
	}

	newResources := []resource.Resource{
		//---Resources
		{Role: "CEO", Type: "human", Name: "Jeph", UserEmail: utils.ValToRef("jeph@jmk21.com")},
		{Role: "AI Bot", Type: "bot", Name: "AI Bot", UserEmail: utils.ValToRef(config.Config.Default.BotUserEmail), ProfileImage: utils.ValToRef("/aibot.jpg"), IsBot: true},
		{Role: "Bartender", Type: "human", Name: "Tifa Lockhart", UserEmail: utils.ValToRef("tifa@jmk21.com"), ProfileImage: utils.ValToRef("/tifa.png"), IsBot: false},
		{Role: "Ex-SOLDIER", Type: "human", Name: "Cloud Strife", UserEmail: utils.ValToRef("cloud@jmk21.com"), ProfileImage: utils.ValToRef("/cloud.png"), IsBot: false},
		{Role: "Florist", Type: "human", Name: "Aerith Gainsborough", UserEmail: utils.ValToRef("aerith@jmk21.com"), ProfileImage: utils.ValToRef("/aerith.png"), IsBot: false},
		{Role: "Avalache", Type: "human", Name: "Barrett Wallace", UserEmail: utils.ValToRef("barrett@jmk21.com"), ProfileImage: utils.ValToRef("/barrett.png"), IsBot: false},
		{Role: "Actress", Type: "human", Name: "Jessie Raspberry", UserEmail: utils.ValToRef("jessie@jmk21.com"), ProfileImage: utils.ValToRef("/jessie.png"), IsBot: false},
		{Role: "Avalache", Type: "human", Name: "Biggs", UserEmail: utils.ValToRef("biggs@jmk21.com"), ProfileImage: utils.ValToRef("/biggs.png"), IsBot: false},
		{Role: "Avalache", Type: "human", Name: "Wedge", UserEmail: utils.ValToRef("wedge@jmk21.com"), ProfileImage: utils.ValToRef("/wedge.png"), IsBot: false},
		{Role: "Pilot", Type: "human", Name: "Cid Highwind", UserEmail: utils.ValToRef("cid@jmk21.com"), ProfileImage: utils.ValToRef("/cid.png"), IsBot: false},
		{Role: "Vampire", Type: "human", Name: "Vincent Valentine", UserEmail: utils.ValToRef("vincent@jmk21.com"), ProfileImage: utils.ValToRef("/vincent.png"), IsBot: false},
		{Role: "Ninja", Type: "human", Name: "Yuffie Kisaragi", UserEmail: utils.ValToRef("yuffie@jmk21.com"), ProfileImage: utils.ValToRef("/yuffie.png"), IsBot: false},
		{Role: "Rat Dog", Type: "human", Name: "Red XIII", UserEmail: utils.ValToRef("red13@jmk21.com"), ProfileImage: utils.ValToRef("/red13.png"), IsBot: false},
		{Role: "SOLDIER", Type: "human", Name: "Zack Fair", UserEmail: utils.ValToRef("zack@jmk21.com"), ProfileImage: utils.ValToRef("/zack.png"), IsBot: false},
	}

	for _, r := range newResources {
		//---this will fail if user already exists, which is fine
		us.GetUser(ctx, *r.UserEmail)
		_, err := rs.CreateResource(ctx, &r)
		if err != nil {
			log.Errorf("error creating resource: %v", err)
		}
	}

	return nil
}
