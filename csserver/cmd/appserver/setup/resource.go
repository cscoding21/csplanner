package setup

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/utils"

	"csserver/internal/services/resource"
	"csserver/internal/services/resource/rtypes"

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
		{Role: "CEO", Type: rtypes.Human, Name: "Jeph", UserEmail: utils.ValToRef("jeph@jmk21.com"), InitialCost: 0.0, AnnualizedCost: 200000, Status: rtypes.Inhouse},
		{Role: "Bartender", Type: rtypes.Human, Name: "Tifa Lockhart", UserEmail: utils.ValToRef("tifa@jmk21.com"), ProfileImage: utils.ValToRef("/tifa.png"), InitialCost: 0.0, AnnualizedCost: 200000, Status: rtypes.Inhouse},
		{Role: "Ex-SOLDIER", Type: rtypes.Human, Name: "Cloud Strife", UserEmail: utils.ValToRef("cloud@jmk21.com"), ProfileImage: utils.ValToRef("/cloud.png"), InitialCost: 0.0, AnnualizedCost: 150000, Status: rtypes.Inhouse},
		{Role: "Florist", Type: rtypes.Human, Name: "Aerith Gainsborough", UserEmail: utils.ValToRef("aerith@jmk21.com"), ProfileImage: utils.ValToRef("/aerith.png"), InitialCost: 0.0, AnnualizedCost: 150000, Status: rtypes.Inhouse},
		{Role: "Avalache", Type: rtypes.Human, Name: "Barrett Wallace", UserEmail: utils.ValToRef("barrett@jmk21.com"), ProfileImage: utils.ValToRef("/barrett.png"), InitialCost: 0.0, AnnualizedCost: 150000, Status: rtypes.Inhouse},
		{Role: "Actress", Type: rtypes.Human, Name: "Jessie Raspberry", UserEmail: utils.ValToRef("jessie@jmk21.com"), ProfileImage: utils.ValToRef("/jessie.png"), InitialCost: 0.0, AnnualizedCost: 150000, Status: rtypes.Inhouse},
		{Role: "Avalache", Type: rtypes.Human, Name: "Biggs", UserEmail: utils.ValToRef("biggs@jmk21.com"), ProfileImage: utils.ValToRef("/biggs.png"), InitialCost: 0.0, AnnualizedCost: 121000, Status: rtypes.Inhouse},
		{Role: "Avalache", Type: rtypes.Human, Name: "Wedge", UserEmail: utils.ValToRef("wedge@jmk21.com"), ProfileImage: utils.ValToRef("/wedge.png"), InitialCost: 0.0, AnnualizedCost: 121000, Status: rtypes.Inhouse},
		{Role: "Pilot", Type: rtypes.Human, Name: "Cid Highwind", UserEmail: utils.ValToRef("cid@jmk21.com"), ProfileImage: utils.ValToRef("/cid.png"), InitialCost: 0.0, AnnualizedCost: 121000, Status: rtypes.Inhouse},
		{Role: "Vampire", Type: rtypes.Human, Name: "Vincent Valentine", UserEmail: utils.ValToRef("vincent@jmk21.com"), ProfileImage: utils.ValToRef("/vincent.png"), InitialCost: 0.0, AnnualizedCost: 121000, Status: rtypes.Inhouse},
		{Role: "Ninja", Type: rtypes.Human, Name: "Yuffie Kisaragi", UserEmail: utils.ValToRef("yuffie@jmk21.com"), ProfileImage: utils.ValToRef("/yuffie.png"), InitialCost: 0.0, AnnualizedCost: 200000, Status: rtypes.Proposed},
		{Role: "Rat Dog", Type: rtypes.Human, Name: "Red XIII", UserEmail: utils.ValToRef("red13@jmk21.com"), ProfileImage: utils.ValToRef("/red13.png"), InitialCost: 0.0, AnnualizedCost: 200000, Status: rtypes.Proposed},
		{Role: "SOLDIER", Type: rtypes.Human, Name: "Zack Fair", UserEmail: utils.ValToRef("zack@jmk21.com"), ProfileImage: utils.ValToRef("/zack.png"), InitialCost: 0.0, AnnualizedCost: 200000, Status: rtypes.Proposed},
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
