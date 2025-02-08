package testobjects

import (
	"csserver/internal/services/organization"
	"csserver/internal/utils"
)

var testOrganization = organization.Organization{
	Name: "Test",
	Defaults: organization.OrganizationDefaults{
		FocusFactor:              utils.ValToRef(5.0),
		CommsCoefficient:         5.0,
		HoursPerWeek:             40,
		DiscountRate:             7.0,
		GenericBlendedHourlyRate: 100.0,
		WorkingHoursPerYear:      2080,
	},
}
