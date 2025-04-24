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

var (
	orgName              = "Jeph Test Org"
	urlKey               = "localhost"
	genesisUserFirstName = "Jeph"
	genesisUserLastName  = "Test"
	genesisUserEmail     = "jeph@jmk21.com"
)

func init() {
	config.InitConfig()
	config.InitLogger()
}

// STANDUP
func TestProvisionNewOrganization(t *testing.T) {
	ctx := getTestContext()
	saasdb := factory.GetSaasDBClient()

	provision.ProvisionNewOrganization(ctx,
		saasdb,
		orgName,
		urlKey,
		genesisUserFirstName,
		genesisUserLastName,
		genesisUserEmail)
}

// TEARDOWN
func TestTeardownOrgInfrastructure(t *testing.T) {
	ctx := getTestContext()
	saasdb := factory.GetSaasDBClient()

	provision.TeardownOrgInfrastructure(ctx, saasdb, orgName)
}

func TestCreateOrgDatabase(t *testing.T) {
	ctx := config.NewContext()
	saasdb := factory.GetSaasDBClient()

	exists := provision.CheckDatabaseExits(ctx, saasdb, orgName)
	fmt.Printf("database %s status - exist = %v\n", orgName, exists)
	var creds config.DatabaseConfig

	if !exists {
		creds, err := provision.CreateOrgDatabase(ctx, saasdb, orgName)
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

	err := provision.AddOrgToMasterDB(ctx, saasdb, orgName)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteMasterDBRecords(t *testing.T) {
	ctx := config.NewContext()
	saasdb := factory.GetSaasDBClient()

	err := provision.DeleteMasterDBRecords(ctx, saasdb, orgName)
	if err != nil {
		t.Error(err)
	}
}

func TestSetOrgProvisioned(t *testing.T) {
	ctx := config.NewContext()
	saasdb := factory.GetSaasDBClient()

	err := provision.SetOrgProvisioned(ctx, saasdb, orgName)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTablesForPlanner(t *testing.T) {
	ctx := config.NewContext()
	creds := provision.GetDBCredsFromName(orgName)
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
	creds := provision.GetDBCredsFromName(orgName)
	os := factory.GetOrganizationService(ctx)
	url := provision.GenerateUrlKeyForOrg(orgName)
	realm := provision.GenerateOrgKey(orgName)

	err := provision.CreateDefaultOrg(ctx, orgName, url, realm, "localhost", creds.Database, os)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateBotUser(t *testing.T) {
	ctx := config.NewContext()
	us := factory.GetIAMAdminService(ctx)
	realm := "jeph_test_org"

	err := provision.CreateBotUser(ctx, realm, us)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateInitialLists(t *testing.T) {
	ctx := config.NewContext()
	ls := factory.GetListService(ctx)

	err := provision.CreateInitialLists(ctx, ls)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateInitialTemplates(t *testing.T) {
	ctx := config.NewContext()
	ts := factory.GetProjectTemplateService(ctx)

	err := provision.CreateInitialTemplates(ctx, ts)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateNewOrgRealm(t *testing.T) {
	ctx := config.NewContext()
	realmName := slug.Make(orgName)

	result, err := provision.CreateNewOrgRealm(ctx, realmName)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Realm '%s' created \n", result)
}

func TestDeleteOrgRealm(t *testing.T) {
	ctx := config.NewContext()
	realmName := slug.Make(orgName)

	err := provision.DeleteOrgRealm(ctx, realmName)
	if err != nil {
		t.Error(err)
	}
}

func TestGetSaaSOrganization(t *testing.T) {
	ctx := config.NewContext()
	url := "localhost"
	saasdb := factory.GetSaasDBClient()

	org, err := provision.GetSaaSOrganization(ctx, saasdb, url)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(org)
}
