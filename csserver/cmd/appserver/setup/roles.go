package setup

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/services/resource"
	"csserver/internal/utils"
	"fmt"
)

func CreateTestRoles(ctx context.Context) error {
	service := factory.GetResourceService()
	roles := []resource.Role{
		{
			ControlFields: common.ControlFields{
				ID: "role:fe",
			},
			Name:        "Frontend Engineer",
			Description: "Creates delightful interfaces",
			HourlyRate:  utils.ValToRef(150.0),
		},
		{
			ControlFields: common.ControlFields{
				ID: "role:be",
			},
			Name:        "Backend Engineer",
			Description: "Creates bug-free backends",
			HourlyRate:  utils.ValToRef(200.0),
		},
		{
			ControlFields: common.ControlFields{
				ID: "role:ui",
			},
			Name:        "UI/UX",
			Description: "Imagines the user experience and interface",
			HourlyRate:  utils.ValToRef(150.0),
		},
		{
			ControlFields: common.ControlFields{
				ID: "role:pm",
			},
			Name:        "Project Manager",
			Description: "Manages the project",
			HourlyRate:  utils.ValToRef(160.0),
		},
		{
			ControlFields: common.ControlFields{
				ID: "role:pdm",
			},
			Name:        "Product Manager",
			Description: "Manages the product",
			HourlyRate:  utils.ValToRef(170.0),
		},
		{
			ControlFields: common.ControlFields{
				ID: "role:sa",
			},
			Name:        "Software Architect",
			Description: "Drives software design",
			HourlyRate:  utils.ValToRef(250.0),
		},
		{
			ControlFields: common.ControlFields{
				ID: "role:te",
			},
			Name:        "Testing Engineer",
			Description: "Ensures product quality",
			HourlyRate:  utils.ValToRef(120.0),
		},
		{
			ControlFields: common.ControlFields{
				ID: "role:mk",
			},
			Name:        "Marketing",
			Description: "Tells theh world",
			HourlyRate:  utils.ValToRef(150.0),
		},
	}

	for _, role := range roles {
		result, err := service.UpsertRole(ctx, role)
		if err != nil {
			fmt.Printf("UpsertRole Primary Error: %s\n", err)
		} else {
			fmt.Printf("UpsertRole Primary: %v\n", result)
		}

	}

	return nil
}
