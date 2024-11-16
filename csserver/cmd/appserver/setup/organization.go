package setup

import (
	"context"

	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/services/organization"

	log "github.com/sirupsen/logrus"
)

// CreateDefaultOrganization creates a default organization
func CreateDefaultOrganization(ctx context.Context) error {
	ts := factory.GetOrganizationService()
	org := &organization.Organization{
		ControlFields: common.ControlFields{
			ID: "organization:default",
		},
		Name: "Default",
		Defaults: organization.OrganizationDefaults{
			FocusFactor:      5.0,
			HoursPerWeek:     40,
			DiscountRate:     7.0,
			CommsCoefficient: 5.0,
		},
	}

	result, err := ts.CreateOrganization(ctx, org)
	if err != nil {
		log.Errorf("CreateOrganization Primary Error: %s\n", err)
		return err
	}

	log.Debugf("CreateOrganization Primary: %v\n", result)

	return nil
}
