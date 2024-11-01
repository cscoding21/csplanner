package setup

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/utils"
	"math/rand"

	"csserver/internal/services/resource"
	"csserver/internal/services/resource/rtypes/resourcestatus"
	"csserver/internal/services/resource/rtypes/resourcetype"

	log "github.com/sirupsen/logrus"
)

func CreateTestResources(ctx context.Context) error {
	//us := factory.GetIAMAdminService()
	rs := factory.GetResourceService()
	exitingResources, _ := rs.FindAllResources(ctx)

	if len(exitingResources.Results) > 0 {
		return nil
	}

	newResources := []resource.Resource{
		//---Resources
		{
			Role:           "CEO",
			Type:           resourcetype.Human,
			Name:           "Jeph",
			UserEmail:      utils.ValToRef("jeph@jmk21.com"),
			InitialCost:    0.0,
			AnnualizedCost: 200000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "security", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Bartender",
			Type:           resourcetype.Human,
			Name:           "Tifa Lockhart",
			UserEmail:      utils.ValToRef("tifa@jmk21.com"),
			ProfileImage:   utils.ValToRef("/tifa.png"),
			InitialCost:    0.0,
			AnnualizedCost: 200000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "marketing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "content-writing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "video-editing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Ex-SOLDIER",
			Type:           resourcetype.Human,
			Name:           "Cloud Strife",
			UserEmail:      utils.ValToRef("cloud@jmk21.com"),
			ProfileImage:   utils.ValToRef("/cloud.png"),
			InitialCost:    0.0,
			AnnualizedCost: 150000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "backend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "database", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},

		{
			Role:           "Florist",
			Type:           resourcetype.Human,
			Name:           "Aerith Gainsborough",
			UserEmail:      utils.ValToRef("aerith@jmk21.com"),
			ProfileImage:   utils.ValToRef("/aerith.png"),
			InitialCost:    0.0,
			AnnualizedCost: 150000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "ui", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "ux", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "frontend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Avalache",
			Type:           resourcetype.Human,
			Name:           "Barret Wallace",
			ProfileImage:   utils.ValToRef("/barrett.png"),
			InitialCost:    0.0,
			AnnualizedCost: 150000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "project-management", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "product-management", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "requirements-gathering", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Actress",
			Type:           resourcetype.Human,
			Name:           "Jessie Raspberry",
			ProfileImage:   utils.ValToRef("/jessie.png"),
			InitialCost:    0.0,
			AnnualizedCost: 150000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "marketing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "content-writing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "video-editing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Avalache",
			Type:           resourcetype.Human,
			Name:           "Biggs",
			ProfileImage:   utils.ValToRef("/biggs.png"),
			InitialCost:    0.0,
			AnnualizedCost: 121000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "project-management", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "requirements-gathering", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Avalache",
			Type:           resourcetype.Human,
			Name:           "Wedge",
			ProfileImage:   utils.ValToRef("/wedge.png"),
			InitialCost:    0.0,
			AnnualizedCost: 121000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "marketing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "content-writing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "frontend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Pilot",
			Type:           resourcetype.Human,
			Name:           "Cid Highwind",
			ProfileImage:   utils.ValToRef("/cid.png"),
			InitialCost:    0.0,
			AnnualizedCost: 121000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "security", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "technical-architecture", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "backend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Vampire",
			Type:           resourcetype.Human,
			Name:           "Vincent Valentine",
			ProfileImage:   utils.ValToRef("/vincent.png"),
			InitialCost:    0.0,
			AnnualizedCost: 121000,
			Status:         resourcestatus.Inhouse,
			Skills: []*resource.Skill{
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "security", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "frontend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Ninja",
			Type:           resourcetype.Human,
			Name:           "Yuffie Kisaragi",
			ProfileImage:   utils.ValToRef("/yuffie.png"),
			InitialCost:    0.0,
			AnnualizedCost: 200000,
			Status:         resourcestatus.Proposed,
			Skills: []*resource.Skill{
				{ID: "business-analysis", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "requirements-gathering", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Rat Dog",
			Type:           resourcetype.Human,
			Name:           "Red XIII",
			ProfileImage:   utils.ValToRef("/red13.png"),
			InitialCost:    0.0,
			AnnualizedCost: 200000,
			Status:         resourcestatus.Proposed,
			Skills: []*resource.Skill{
				{ID: "business-analysis", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "requirements-gathering", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "SOLDIER",
			Type:           resourcetype.Human,
			Name:           "Zack Fair",
			ProfileImage:   utils.ValToRef("/zack.png"),
			InitialCost:    0.0,
			AnnualizedCost: 200000,
			Status:         resourcestatus.Proposed,
			Skills: []*resource.Skill{
				{ID: "technical-architecture", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "backend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "database", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			Role:           "Sword",
			Type:           resourcetype.Equipment,
			Name:           "Buster Sword",
			InitialCost:    10000,
			AnnualizedCost: 0,
			Status:         resourcestatus.Proposed,
		},
		{
			Role:           "Cloud Hosting",
			Type:           resourcetype.Software,
			Name:           "Linode Cloud Hosting",
			InitialCost:    0,
			AnnualizedCost: 100000,
			Status:         resourcestatus.Proposed,
		},
	}

	for _, r := range newResources {
		//---this will fail if user already exists, which is fine
		if r.UserEmail != nil {
			//us.GetUser(ctx, *r.UserEmail)
		}
		_, err := rs.CreateResource(ctx, &r)
		if err != nil {
			log.Errorf("error creating resource: %v", err)
		}
	}

	return nil
}
