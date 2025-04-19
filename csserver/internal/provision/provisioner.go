package provision

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/providers/postgres"
	"csserver/internal/services/iam/appuser"
	"csserver/internal/services/iam/auth"
	"csserver/internal/services/list"
	"csserver/internal/services/organization"
	"csserver/internal/services/projecttemplate"
	"csserver/internal/utils"
	"errors"
	"fmt"
	"strings"

	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5/pgxpool"

	log "github.com/sirupsen/logrus"
)

func ProvisionNewOrganization(
	ctx context.Context,
	db *pgxpool.Pool,
	name string) error {

	existing := CheckDatabaseExits(ctx, db, name)
	if !existing {
		_, err := CreateOrgDatabase(ctx, db, name)
		if err != nil {
			log.Fatal(err)
		}
	}

	orgDBCreds := GetDBCredsFromName(name)
	orgDBCreds.Host = db.Config().ConnConfig.Host
	orgDBCreds.Port = int(db.Config().ConnConfig.Port)
	orgDBCreds.User = db.Config().ConnConfig.User
	orgDBCreds.Password = db.Config().ConnConfig.Password

	orgDBClient, err := postgres.GetDBFromConfig(ctx, orgDBCreds)
	if err != nil {
		log.Fatal(err)
	}

	errs := CreateTablesForPlanner(ctx, orgDBClient)
	if errs != nil {
		log.Error(errs)
	}

	CreateNewOrgRealm(ctx, name)

	pubSub, _ := factory.GetPubSubClient()
	gk := getKeycloakClient()
	appUserService := appuser.NewAppuserService(orgDBClient, pubSub)
	us := auth.NewIAMAdminService(gk, pubSub, orgDBCreds.Database, config.Config.Security.KeycloakAdminUser, config.Config.Security.KeycloakAdminPass, *appUserService)

	err = CreateBotUser(ctx, &us)
	if err != nil {
		log.Fatal(err)
	}

	/*
		check for existing name and provision status in master DB table
		- check for existing db
			-  if yes, exit

		create db for new org
		create realm for new org

		run setup scripts/migrations
		- create bot user
		- create default organization
		- create initial lists
			- skills (no values)
			- funding sources (internal/external)
			- value categories ()
		- create project templates
			- simple
			- complex

	*/

	return nil
}

// TeardownOrgInfrastructure remove all elements of a SaaS org
func TeardownOrgInfrastructure(
	ctx context.Context,
	db *pgxpool.Pool,
	name string) {

	//--- delete database and perform all cleanup
	err := dropOrgDatabase(ctx, db, name)
	if err != nil {
		log.Error(err)
	}

	//--- delete realm
	err = DeleteOrgRealm(ctx, name)
	if err != nil {
		log.Error(err)
	}

	//--- delete master db records
	err = DeleteMasterDBRecords(ctx, db, name)
	if err != nil {
		log.Error(err)
	}
}

func dropOrgDatabase(
	ctx context.Context,
	db *pgxpool.Pool,
	name string) error {

	newDBCreds := GetDBCredsFromName(name)

	// 	var dropDatabaseSQL = `DROP DATABASE %s WITH(force);`
	// var revokeTablePrivligesDatabaseSQL = `REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA public FROM %s;`
	// var revokeSchemaPrivligesDatabaseSQL = `REVOKE USAGE ON SCHEMA public FROM %s;`
	// var dropDatabaseUserSQL = `DROP USER %s;`

	sql := fmt.Sprintf(dropDatabaseSQL, newDBCreds.Database)
	err := errors.Join(postgres.Exec(ctx, db, sql))
	if err != nil {
		log.Error(err)
		return err
	}

	sql = fmt.Sprintf(revokeTablePrivligesDatabaseSQL, newDBCreds.Database)
	err = errors.Join(postgres.Exec(ctx, db, sql))
	if err != nil {
		log.Error(err)
		return err
	}

	sql = fmt.Sprintf(revokeSchemaPrivligesDatabaseSQL, newDBCreds.Database)
	err = errors.Join(postgres.Exec(ctx, db, sql))
	if err != nil {
		log.Error(err)
		return err
	}

	sql = fmt.Sprintf(dropDatabaseUserSQL, newDBCreds.Database)
	err = errors.Join(postgres.Exec(ctx, db, sql))
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func DeleteMasterDBRecords(
	ctx context.Context,
	db *pgxpool.Pool,
	name string) error {

	id := getOrgIDFromName(name)

	err := errors.Join(postgres.Exec(ctx, db, deleteOrgLicensesSQL, id))
	if err != nil {
		log.Error(err)
		return err
	}

	err = errors.Join(postgres.Exec(ctx, db, deleteOrgSQL, id))
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// CheckDatabaseExits return true if the proposed database exists
func CheckDatabaseExits(
	ctx context.Context,
	db *pgxpool.Pool,
	name string) bool {

	sql := checkDatabaseExistSQL
	count, err := postgres.GetScalar[int](ctx, db, sql, name)
	if err != nil {
		return false
	}

	return count != nil && *count == 1
}

// CreateOrgDatabase create a new database for a give org
func CreateOrgDatabase(
	ctx context.Context,
	db *pgxpool.Pool,
	name string) (config.DatabaseConfig, error) {

	newDBCreds := GetDBCredsFromName(name)

	//---create DB
	sql := fmt.Sprintf(newDatabaseSQL, newDBCreds.Database)
	err := postgres.Exec(ctx, db, sql)
	if err != nil {
		log.Error(err)
		return newDBCreds, err
	}

	//---create user
	sql = fmt.Sprintf(newDBUserSQL, newDBCreds.User, newDBCreds.Password)
	err = errors.Join(postgres.Exec(ctx, db, sql))
	if err != nil {
		log.Error(err)
	}

	//---grant table privilidges
	sql = fmt.Sprintf(grantOnDBSQL, newDBCreds.Database, newDBCreds.User)
	err = errors.Join(postgres.Exec(ctx, db, sql))
	if err != nil {
		log.Error(err)
	}

	//---grant schema privilidges
	// sql = fmt.Sprintf(grandOnSchemaSQL, newDBCreds.User)
	// err = errors.Join(postgres.Exec(ctx, db, sql))
	// if err != nil {
	// 	log.Error(err)
	// }

	return newDBCreds, err
}

// AddOrgToMasterDB add the organization record to the master DB list
func AddOrgToMasterDB(
	ctx context.Context,
	db *pgxpool.Pool,
	name string) error {

	//---TODO: create tokenize function
	tokenizedName := slug.Make(name)
	id := getOrgIDFromName(name)
	opUser := "test:bot"

	sql := newMasterOrgRecordSQL
	return postgres.Exec(ctx, db, sql, id, name, tokenizedName, false, "localhost", opUser, opUser)
}

func getOrgIDFromName(name string) string {
	tokenizedName := slug.Make(name)
	id := fmt.Sprintf("organization:%s", tokenizedName)

	return id
}

// SetOrgProvisioned set the orgs provisioned flag to true.  This should be called when all steps have been completed
func SetOrgProvisioned(
	ctx context.Context,
	db *pgxpool.Pool,
	name string) error {
	opUser := "test:bot"

	orgID := fmt.Sprintf("organization:%s", slug.Make(name))
	sql := setProvisionedCompleteSQL
	return postgres.Exec(ctx, db, sql, opUser, orgID)
}

// CreateTablesForPlanner generaete all tables required for a csPlanner instance DB
func CreateTablesForPlanner(
	ctx context.Context,
	db *pgxpool.Pool) error {

	tables := []string{"list", "organization", "project", "projecttemplate", "resource", "role", "appuser", "comment", "reaction"}

	var err error

	for _, t := range tables {
		fmt.Println(t)

		sql := fmt.Sprintf(createCSPlannerTableSQL, t, t, t, t, t, t, t, t, t, t, t, t, t)
		errors.Join(err, postgres.Exec(ctx, db, sql))
	}

	return err
}

// CreateDefaultOrg create the default organization for a new csPlanner instance
func CreateDefaultOrg(
	ctx context.Context,
	name string,
	url string,
	orgService *organization.OrganizationService) error {
	org := organization.Organization{
		ControlFields: common.ControlFields{
			ID: organization.DefaultOrganizationID,
		},
		Name: name,
		URL:  url,
		Defaults: organization.OrganizationDefaults{
			FocusFactor:              utils.ValToRef(5.0),
			HoursPerWeek:             40,
			DiscountRate:             7.0,
			CommsCoefficient:         5.0,
			GenericBlendedHourlyRate: 100.0,
			WorkingHoursPerYear:      2080,
		},
	}

	result, err := orgService.CreateOrganization(ctx, org)
	if err != nil {
		log.Errorf("CreateOrganization Error: %s\n", err)
		return err
	}

	log.Debugf("CreateOrganization: %v\n", result)

	return nil
}

// CreateBotUser create the bot user account for a csPlanner instance
func CreateBotUser(
	ctx context.Context,
	userService *auth.IAMAdminService) error {

	botUser := appuser.Appuser{
		ControlFields: common.ControlFields{
			ID: "user:bot",
		},
		Email:           config.Config.Default.BotUserEmail,
		FirstName:       "AI",
		LastName:        "Bot",
		Password:        config.Config.Default.BotPassword,
		ConfirmPassword: config.Config.Default.BotPassword,
		ProfileImage:    "/aibot.jpg",
	}

	_, err := userService.CreateUser(ctx, &botUser)
	if err != nil {
		log.Errorf("error creating user: %v", err)
		return err
	}

	return nil
}

// CreateInitialLists create all of the required list records needed for csPlanner
func CreateInitialLists(
	ctx context.Context,
	listService *list.ListService) error {

	existingList, err := listService.GetList(ctx, list.ListNameSkills)
	if existingList == nil {
		log.Errorf("Error getting skills list: %v", err)

		ur, err := listService.CreateList(ctx, skillsList)
		if err != nil {
			log.Errorf("Error creating skills list: %v", err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	}

	existingList, err = listService.GetList(ctx, list.ListNameFundingSource)
	if existingList == nil {
		log.Errorf("Error getting funding source list: %v", err)

		ur, err := listService.CreateList(ctx, fundingSourceList)
		if err != nil {
			log.Errorf("Error creating funding source list: %v", err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	}

	existingList, err = listService.GetList(ctx, list.ListNameValueCategory)
	if existingList == nil {
		log.Errorf("Error getting value category list: %v", err)

		ur, err := listService.CreateList(ctx, valueCategoryList)
		if err != nil {
			log.Errorf("Error creating value category list: %v", err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	}

	return nil
}

// CreateInitialTemplates create a default set of project templates for csPlanner
func CreateInitialTemplates(
	ctx context.Context,
	templateService *projecttemplate.ProjecttemplateService) error {

	result, err := templateService.SaveTemplate(ctx, *simpleTemplate)
	if err != nil {
		fmt.Printf("CreateInitialTemplates Simple Error: %s\n", err)
	} else {
		fmt.Printf("CreateInitialTemplates Simple: %v\n", result)
	}

	result, err = templateService.SaveTemplate(ctx, *template)
	if err != nil {
		fmt.Printf("CreateInitialTemplates Tempate Error: %s\n", err)
	} else {
		fmt.Printf("CreateInitialTemplates Template: %v\n", result)
	}

	return nil
}

// GetDBCredsFromName return a DB creds object based on the
func GetDBCredsFromName(name string) config.DatabaseConfig {
	name = slug.Make(name)
	name = strings.Replace(name, "-", "_", -1)
	return config.DatabaseConfig{
		Host:     config.Config.MasterDB.Host,
		Port:     config.Config.MasterDB.Port,
		Database: fmt.Sprintf("csp_%s", name),
		User:     fmt.Sprintf("csp_%s_user", name),
		Password: utils.GeneratePassword(),
	}
}
