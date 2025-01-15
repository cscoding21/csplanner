package testobjects

import "csserver/internal/services/organization"

var testOrganization = organization.Organization{
	Name: "Test",
	Defaults: organization.OrganizationDefaults{
		FocusFactor:      5.0,
		CommsCoefficient: 5.0,
		HoursPerWeek:     40,
		DiscountRate:     7.0,
	},
}
