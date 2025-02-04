package setup

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/common"
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
			ControlFields: common.ControlFields{
				ID: "resource:jeph",
			},
			RoleID:                utils.ValToRef("role:sa"),
			Type:                  resourcetype.Human,
			Name:                  "Jeph",
			UserEmail:             utils.ValToRef("jeph@jmk21.com"),
			InitialCost:           0.0,
			AnnualizedCost:        200000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 32,
			Skills: []*resource.Skill{
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "security", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:tifa",
			},
			RoleID:                utils.ValToRef("role:fe"),
			Type:                  resourcetype.Human,
			Name:                  "Tifa Lockhart",
			UserEmail:             utils.ValToRef("tifa@jmk21.com"),
			ProfileImage:          utils.ValToRef("/tifa.png"),
			InitialCost:           0.0,
			AnnualizedCost:        200000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 32,
			Skills: []*resource.Skill{
				{ID: "marketing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "content-writing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "video-editing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:cloud",
			},
			RoleID:                utils.ValToRef("role:be"),
			Type:                  resourcetype.Human,
			Name:                  "Cloud Strife",
			UserEmail:             utils.ValToRef("cloud@jmk21.com"),
			ProfileImage:          utils.ValToRef("/cloud.png"),
			InitialCost:           0.0,
			AnnualizedCost:        150000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "backend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "database", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},

		{
			ControlFields: common.ControlFields{
				ID: "resource:aerith",
			},
			RoleID:                utils.ValToRef("role:ui"),
			Type:                  resourcetype.Human,
			Name:                  "Aerith Gainsborough",
			UserEmail:             utils.ValToRef("aerith@jmk21.com"),
			ProfileImage:          utils.ValToRef("/aerith.png"),
			InitialCost:           0.0,
			AnnualizedCost:        150000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: "ui", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "ux", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "frontend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:barret",
			},
			RoleID:                utils.ValToRef("role:pm"),
			Type:                  resourcetype.Human,
			Name:                  "Barret Wallace",
			ProfileImage:          utils.ValToRef("/barrett.png"),
			InitialCost:           0.0,
			AnnualizedCost:        150000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: "project-management", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "product-management", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "requirements-gathering", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "communications", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "ui", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:jessie",
			},
			RoleID:                utils.ValToRef("role:mk"),
			Type:                  resourcetype.Human,
			Name:                  "Jessie Raspberry",
			ProfileImage:          utils.ValToRef("/jessie.png"),
			InitialCost:           0.0,
			AnnualizedCost:        150000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: "marketing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "content-writing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "video-editing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:biggs",
			},
			RoleID:                utils.ValToRef("role:be"),
			Type:                  resourcetype.Human,
			Name:                  "Biggs",
			ProfileImage:          utils.ValToRef("/biggs.png"),
			InitialCost:           0.0,
			AnnualizedCost:        121000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 32,
			Skills: []*resource.Skill{
				{ID: "project-management", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "requirements-gathering", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:wedge",
			},
			RoleID:                utils.ValToRef("role:mk"),
			Type:                  resourcetype.Human,
			Name:                  "Wedge",
			ProfileImage:          utils.ValToRef("/wedge.png"),
			InitialCost:           0.0,
			AnnualizedCost:        121000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: "marketing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "content-writing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "frontend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "technical-writing", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:cid",
			},
			RoleID:                utils.ValToRef("role:ta"),
			Type:                  resourcetype.Human,
			Name:                  "Cid Highwind",
			ProfileImage:          utils.ValToRef("/cid.png"),
			InitialCost:           0.0,
			AnnualizedCost:        121000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 20,
			Skills: []*resource.Skill{
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "security", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "technical-architecture", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "backend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:vincent",
			},
			RoleID:                utils.ValToRef("role:te"),
			Type:                  resourcetype.Human,
			Name:                  "Vincent Valentine",
			ProfileImage:          utils.ValToRef("/vincent.png"),
			InitialCost:           0.0,
			AnnualizedCost:        121000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "security", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "frontend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:yuffie",
			},
			RoleID:                utils.ValToRef("role:pdm"),
			Type:                  resourcetype.Human,
			Name:                  "Yuffie Kisaragi",
			ProfileImage:          utils.ValToRef("/yuffie.png"),
			InitialCost:           0.0,
			AnnualizedCost:        200000,
			Status:                resourcestatus.Proposed,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: "business-analysis", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "requirements-gathering", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:red",
			},
			RoleID:                utils.ValToRef("role:pdm"),
			Type:                  resourcetype.Human,
			Name:                  "Red XIII",
			ProfileImage:          utils.ValToRef("/red13.png"),
			InitialCost:           0.0,
			AnnualizedCost:        200000,
			Status:                resourcestatus.Proposed,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: "business-analysis", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "requirements-gathering", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:zack",
			},
			RoleID:                utils.ValToRef("role:ta"),
			Type:                  resourcetype.Human,
			Name:                  "Zack Fair",
			ProfileImage:          utils.ValToRef("/zack.png"),
			InitialCost:           0.0,
			AnnualizedCost:        200000,
			Status:                resourcestatus.Proposed,
			AvailableHoursPerWeek: 32,
			Skills: []*resource.Skill{
				{ID: "technical-architecture", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "devops", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "backend", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
				{ID: "database", Proficiency: utils.ValToRef(1.0 + float64(rand.Intn(3)))},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:buster",
			},
			Type:           resourcetype.Equipment,
			Name:           "Buster Sword",
			InitialCost:    10000,
			AnnualizedCost: 0,
			Status:         resourcestatus.Proposed,
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:linode",
			},
			Type:           resourcetype.Software,
			Name:           "Linode Cloud Hosting",
			InitialCost:    0,
			AnnualizedCost: 100000,
			Status:         resourcestatus.Proposed,
		},
	}

	for _, r := range newResources {
		_, err := rs.CreateResource(ctx, &r)
		if err != nil {
			log.Errorf("error creating resource: %v", err)
		}
	}

	return nil
}
