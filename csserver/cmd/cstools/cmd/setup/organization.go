package setup

import (
	"context"

	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/services/organization"
	"csserver/internal/utils"

	log "github.com/sirupsen/logrus"
)

// CreateDefaultOrganization creates a default organization
func CreateDefaultOrganization(ctx context.Context) error {
	ts := factory.GetOrganizationService(ctx)
	org := organization.Organization{
		ControlFields: common.ControlFields{
			ID: organization.DefaultOrganizationID,
		},
		Name:     "Jeph Test Org",
		URLKey:   "localhost",
		DBHost:   "localhost",
		Realm:    "jeph_test_org",
		Database: "csp_jeph_test_org",
		Defaults: organization.OrganizationDefaults{
			FocusFactor:              utils.ValToRef(5.0),
			HoursPerWeek:             40,
			DiscountRate:             7.0,
			CommsCoefficient:         5.0,
			GenericBlendedHourlyRate: 100.0,
			WorkingHoursPerYear:      2080,
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
