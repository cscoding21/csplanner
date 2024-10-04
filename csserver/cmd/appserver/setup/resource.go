package setup

import (
	"context"
	"csserver/internal/appserv/factory"
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
		{Role: "Bartender", Type: "human", Name: "Tifa Lockhart", UserEmail: utils.ValToRef("tifa@jmk21.com"), ProfileImage: utils.ValToRef("/tifa.png")},
		{Role: "Ex-SOLDIER", Type: "human", Name: "Cloud Strife", UserEmail: utils.ValToRef("cloud@jmk21.com"), ProfileImage: utils.ValToRef("/cloud.png")},
		{Role: "Florist", Type: "human", Name: "Aerith Gainsborough", UserEmail: utils.ValToRef("aerith@jmk21.com"), ProfileImage: utils.ValToRef("/aerith.png")},
		{Role: "Avalache", Type: "human", Name: "Barrett Wallace", UserEmail: utils.ValToRef("barrett@jmk21.com"), ProfileImage: utils.ValToRef("/barrett.png")},
		{Role: "Actress", Type: "human", Name: "Jessie Raspberry", UserEmail: utils.ValToRef("jessie@jmk21.com"), ProfileImage: utils.ValToRef("/jessie.png")},
		{Role: "Avalache", Type: "human", Name: "Biggs", UserEmail: utils.ValToRef("biggs@jmk21.com"), ProfileImage: utils.ValToRef("/biggs.png")},
		{Role: "Avalache", Type: "human", Name: "Wedge", UserEmail: utils.ValToRef("wedge@jmk21.com"), ProfileImage: utils.ValToRef("/wedge.png")},
		{Role: "Pilot", Type: "human", Name: "Cid Highwind", UserEmail: utils.ValToRef("cid@jmk21.com"), ProfileImage: utils.ValToRef("/cid.png")},
		{Role: "Vampire", Type: "human", Name: "Vincent Valentine", UserEmail: utils.ValToRef("vincent@jmk21.com"), ProfileImage: utils.ValToRef("/vincent.png")},
		{Role: "Ninja", Type: "human", Name: "Yuffie Kisaragi", UserEmail: utils.ValToRef("yuffie@jmk21.com"), ProfileImage: utils.ValToRef("/yuffie.png")},
		{Role: "Rat Dog", Type: "human", Name: "Red XIII", UserEmail: utils.ValToRef("red13@jmk21.com"), ProfileImage: utils.ValToRef("/red13.png")},
		{Role: "SOLDIER", Type: "human", Name: "Zack Fair", UserEmail: utils.ValToRef("zack@jmk21.com"), ProfileImage: utils.ValToRef("/zack.png")},
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
