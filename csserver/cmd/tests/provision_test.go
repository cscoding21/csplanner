package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"csserver/internal/providers/postgres"
	"csserver/internal/provision"
	"fmt"
	"testing"

	"github.com/gosimple/slug"
)

var name = "Jeph Test Org"

func init() {
	config.InitConfig()
	config.InitLogger()
}

// STANDUP
func TestProvisionNewOrganization(t *testing.T) {
	ctx := getTestContext()
	saasdb := factory.GetSaasDBClient()

	provision.ProvisionNewOrganization(ctx, saasdb, name)
}

// TEARDOWN
func TestTeardownOrgInfrastructure(t *testing.T) {
	ctx := getTestContext()
	saasdb := factory.GetSaasDBClient()

	provision.TeardownOrgInfrastructure(ctx, saasdb, name)
}

func TestCreateOrgDatabase(t *testing.T) {
	ctx := config.NewContext()
	saasdb := factory.GetSaasDBClient()

	exists := provision.CheckDatabaseExits(ctx, saasdb, name)
	fmt.Printf("database %s status - exist = %v\n", name, exists)
	var creds config.DatabaseConfig

	if !exists {
		creds, err := provision.CreateOrgDatabase(ctx, saasdb, name)
		if err != nil {
			t.Error(err)
			return
		}

		fmt.Println(creds)
	}

	fmt.Println(creds.Password)
}

func TestAddOrgToMasterDB(t *testing.T) {
	ctx := config.NewContext()
	saasdb := factory.GetSaasDBClient()

	err := provision.AddOrgToMasterDB(ctx, saasdb, name)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteMasterDBRecords(t *testing.T) {
	ctx := config.NewContext()
	saasdb := factory.GetSaasDBClient()

	err := provision.DeleteMasterDBRecords(ctx, saasdb, name)
	if err != nil {
		t.Error(err)
	}
}

func TestSetOrgProvisioned(t *testing.T) {
	ctx := config.NewContext()
	saasdb := factory.GetSaasDBClient()

	err := provision.SetOrgProvisioned(ctx, saasdb, name)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTablesForPlanner(t *testing.T) {
	ctx := config.NewContext()
	creds := provision.GetDBCredsFromName(name)
	creds.User = "postgres"
	creds.Password = "postgres"
	db, err := postgres.GetDBFromConfig(ctx, creds)
	if err != nil {
		t.Error(err)
		return
	}

	err = provision.CreateTablesForPlanner(ctx, db)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateDefaultOrg(t *testing.T) {
	ctx := config.NewContext()
	//creds := provision.GetDBCredsFromName(name)
	os := factory.GetOrganizationService()
	url := slug.Make(name)

	err := provision.CreateDefaultOrg(ctx, name, url, os)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateBotUser(t *testing.T) {
	ctx := config.NewContext()
	us := factory.GetIAMAdminService()

	err := provision.CreateBotUser(ctx, us)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateInitialLists(t *testing.T) {
	ctx := config.NewContext()
	ls := factory.GetListService()

	err := provision.CreateInitialLists(ctx, ls)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateInitialTemplates(t *testing.T) {
	ctx := config.NewContext()
	ts := factory.GetProjectTemplateService()

	err := provision.CreateInitialTemplates(ctx, ts)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateNewOrgRealm(t *testing.T) {
	ctx := config.NewContext()
	realmName := slug.Make(name)

	result, err := provision.CreateNewOrgRealm(ctx, realmName)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Realm '%s' created \n", result)
}

func TestDeleteOrgRealm(t *testing.T) {
	ctx := config.NewContext()
	realmName := slug.Make(name)

	err := provision.DeleteOrgRealm(ctx, realmName)
	if err != nil {
		t.Error(err)
	}
}
