package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/appserv/orgmap"
	"fmt"
	"testing"
)

func TestGetOrganization(t *testing.T) {
	ctx := getTestContext()
	os := factory.GetOrganizationService(ctx)

	om, err := orgmap.GetSaaSOrg(ctx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(om.DB.Config().ConnConfig.Host)

	o, err := os.GetDefaultOrganization(ctx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(o)
}
