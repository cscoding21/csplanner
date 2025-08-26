package setup

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/utils"

	"csserver/internal/services/resource"
	"csserver/internal/services/resource/rtypes/resourcestatus"
	"csserver/internal/services/resource/rtypes/resourcetype"

	log "github.com/sirupsen/logrus"
)

func CreateTestResources(ctx context.Context) error {
	//us := factory.GetIAMAdminService()
	rs := factory.GetResourceService(ctx)
	org, err := factory.GetDefaultOrganization(ctx)
	if err != nil {
		return err
	}

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
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 32,
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:tori",
			},
			RoleID:                utils.ValToRef("role:fe"),
			Type:                  resourcetype.Human,
			Name:                  "Tori L",
			UserEmail:             utils.ValToRef("tori@jmk21.com"),
			ProfileImage:          utils.ValToRef("/team-01.png"),
			InitialCost:           0.0,
			AnnualizedCost:        200000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 32,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:marketing", Proficiency: utils.ValToRef(3.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:content-writing", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:video-editing", Proficiency: utils.ValToRef(1.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:chris",
			},
			RoleID:                utils.ValToRef("role:be"),
			Type:                  resourcetype.Human,
			Name:                  "Chris S",
			UserEmail:             utils.ValToRef("chris@jmk21.com"),
			ProfileImage:          utils.ValToRef("/team-02.png"),
			InitialCost:           0.0,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:devops", Proficiency: utils.ValToRef(1.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:backend", Proficiency: utils.ValToRef(3.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:database", Proficiency: utils.ValToRef(2.0)},
			},
		},

		{
			ControlFields: common.ControlFields{
				ID: "resource:adam",
			},
			RoleID:                utils.ValToRef("role:ui"),
			Type:                  resourcetype.Human,
			Name:                  "Adam G",
			UserEmail:             utils.ValToRef("adam@jmk21.com"),
			ProfileImage:          utils.ValToRef("/team-03.png"),
			InitialCost:           0.0,
			AnnualizedCost:        150000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:ui", Proficiency: utils.ValToRef(3.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:ux", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:frontend", Proficiency: utils.ValToRef(1.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:barry",
			},
			RoleID:                utils.ValToRef("role:pm"),
			Type:                  resourcetype.Human,
			Name:                  "Barry W",
			ProfileImage:          utils.ValToRef("/team-04.png"),
			InitialCost:           0.0,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:project-management", Proficiency: utils.ValToRef(3.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:product-management", Proficiency: utils.ValToRef(1.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:requirements-gathering", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:communications", Proficiency: utils.ValToRef(3.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:ui", Proficiency: utils.ValToRef(1.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:jerry",
			},
			RoleID:                utils.ValToRef("role:mk"),
			Type:                  resourcetype.Human,
			Name:                  "Jerry R",
			ProfileImage:          utils.ValToRef("/team-05.png"),
			InitialCost:           0.0,
			AnnualizedCost:        121000,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:marketing", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:content-writing", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:video-editing", Proficiency: utils.ValToRef(1.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:brenda",
			},
			RoleID:                utils.ValToRef("role:pdm"),
			Type:                  resourcetype.Human,
			Name:                  "Brenda",
			ProfileImage:          utils.ValToRef("/team-06.png"),
			InitialCost:           0.0,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 32,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:project-management", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:requirements-gathering", Proficiency: utils.ValToRef(2.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:wendy",
			},
			RoleID:                utils.ValToRef("role:mk"),
			Type:                  resourcetype.Human,
			Name:                  "Wendy H",
			ProfileImage:          utils.ValToRef("/team-07.png"),
			InitialCost:           0.0,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:marketing", Proficiency: utils.ValToRef(3.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:content-writing", Proficiency: utils.ValToRef(1.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:frontend", Proficiency: utils.ValToRef(1.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:technical-writing", Proficiency: utils.ValToRef(1.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:celine",
			},
			RoleID:                utils.ValToRef("role:sa"),
			Type:                  resourcetype.Human,
			Name:                  "Celine H",
			ProfileImage:          utils.ValToRef("/team-08.png"),
			InitialCost:           0.0,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 20,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:devops", Proficiency: utils.ValToRef(3.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:security", Proficiency: utils.ValToRef(1.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:technical-architecture", Proficiency: utils.ValToRef(3.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:backend", Proficiency: utils.ValToRef(2.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:vinnie",
			},
			RoleID:                utils.ValToRef("role:te"),
			Type:                  resourcetype.Human,
			Name:                  "Vinnie V",
			ProfileImage:          utils.ValToRef("/team-09.png"),
			InitialCost:           0.0,
			Status:                resourcestatus.Inhouse,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:security", Proficiency: utils.ValToRef(1.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:frontend", Proficiency: utils.ValToRef(2.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:yuri",
			},
			RoleID:                utils.ValToRef("role:pdm"),
			Type:                  resourcetype.Human,
			Name:                  "Yuri K",
			ProfileImage:          utils.ValToRef("/team-10.png"),
			InitialCost:           0.0,
			AnnualizedCost:        140000,
			Status:                resourcestatus.Proposed,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:business-analysis", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:requirements-gathering", Proficiency: utils.ValToRef(2.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:red",
			},
			RoleID:                utils.ValToRef("role:pdm"),
			Type:                  resourcetype.Human,
			Name:                  "Reb N",
			ProfileImage:          utils.ValToRef("/team-11.png"),
			InitialCost:           0.0,
			Status:                resourcestatus.Proposed,
			AvailableHoursPerWeek: 40,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:business-analysis", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:requirements-gathering", Proficiency: utils.ValToRef(2.0)},
			},
		},
		{
			ControlFields: common.ControlFields{
				ID: "resource:zelda",
			},
			RoleID:                utils.ValToRef("role:sa"),
			Type:                  resourcetype.Human,
			Name:                  "Zelda F",
			ProfileImage:          utils.ValToRef("/team-12.png"),
			InitialCost:           0.0,
			Status:                resourcestatus.Proposed,
			AvailableHoursPerWeek: 32,
			Skills: []*resource.Skill{
				{ID: utils.GenerateBase64UUID(), SkillID: "td:technical-architecture", Proficiency: utils.ValToRef(3.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:devops", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:backend", Proficiency: utils.ValToRef(2.0)},
				{ID: utils.GenerateBase64UUID(), SkillID: "td:database", Proficiency: utils.ValToRef(1.0)},
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
		_, err := rs.SaveResource(ctx, r, *org)
		if err != nil {
			log.Errorf("error creating resource: %v", err)
		}
	}

	return nil
}
