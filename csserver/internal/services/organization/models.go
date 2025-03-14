//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package organization

import (
	"csserver/internal/common"
)

type Organization struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	Name     string               `json:"name"`
	URL      string               `json:"url"`
	Defaults OrganizationDefaults `json:"organization_defaults"`
}

type OrganizationDefaults struct {
	DiscountRate             float64  `json:"discount_rate"`
	HoursPerWeek             int      `json:"hours_per_week"`
	FocusFactor              *float64 `json:"focus_factor"`
	CommsCoefficient         float64  `json:"comms_coefficient"`
	GenericBlendedHourlyRate int      `json:"generic_blended_hourly_rate"`
	WorkingHoursPerYear      int      `json:"working_hours_per_year"`
}
