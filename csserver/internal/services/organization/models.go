//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package organization

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/config"
	"strings"
	"time"
)

type Organization struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	Name     string               `json:"name"`
	URLKey   string               `json:"url_key"`
	DBHost   string               `json:"db_host"`
	Database string               `json:"database"`
	Realm    string               `json:"realm"`
	Defaults OrganizationDefaults `json:"organization_defaults"`
	Logo     string               `json:"logo"`
	Licenses []License            `json:"licenses"`
}

type OrganizationDefaults struct {
	DiscountRate             float64  `json:"discount_rate"`
	HoursPerWeek             int      `json:"hours_per_week"`
	FocusFactor              *float64 `json:"focus_factor"`
	CommsCoefficient         float64  `json:"comms_coefficient"`
	GenericBlendedHourlyRate int      `json:"generic_blended_hourly_rate"`
	WorkingHoursPerYear      int      `json:"working_hours_per_year"`
}

type License struct {
	ID          string     `json:"id,omitempty"`
	OrgID       string     `json:"org_id"`
	App         string     `json:"app"`
	Email       string     `json:"email"`
	ExpiresOn   time.Time  `json:"expires_on"`
	CancelledOn *time.Time `json:"cancelled_on"`
	AutoRenew   bool       `json:"auto_renew"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   string     `json:"updated_by"`
}

type SaaSInfo struct {
	Org      *SaaSOrg
	Licenses *[]License
}
type SaaSOrg struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	DBHost        string    `json:"db_host"`
	Database      string    `json:"database"`
	UrlKey        string    `json:"url_key"`
	Realm         string    `json:"realm"`
	IsProvisioned bool      `json:"is_provisioned"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedAt     time.Time `json:"updated_at"`
	UpdatedBy     string    `json:"updated_by"`
}

func (so *SaaSInfo) HasValidLicense(ctx context.Context) bool {
	if so.Licenses == nil || len(*so.Licenses) == 0 {
		return false
	}

	user := config.GetUserEmailFromContext(ctx)

	for _, l := range *so.Licenses {
		if strings.EqualFold(user, l.Email) {
			return true
		}
	}

	return false
}
