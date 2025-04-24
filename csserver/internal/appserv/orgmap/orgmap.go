package orgmap

import (
	"context"
	"csserver/internal/config"
	"csserver/internal/providers/postgres"
	"csserver/internal/services/organization"
	"csserver/internal/utils"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"

	log "github.com/sirupsen/logrus"
)

var (
	orgmap       map[string]*OrgResources
	saasDBClient *pgxpool.Pool
	lock         = &sync.Mutex{}
)

type OrgResources struct {
	Info *organization.SaaSInfo
	DB   *pgxpool.Pool
}

func GetSaaSOrg(ctx context.Context) (*OrgResources, error) {
	if orgmap == nil {
		orgmap = make(map[string]*OrgResources)
	}

	urlKey := config.GetOrgUrlKeyFromContext(ctx)
	if len(urlKey) == 0 {
		log.Errorf("GetOrgUrlKeyFromContext: no org key found in context")
	}

	outOrg, ok := orgmap[urlKey]
	if ok {
		log.Debugf("returning org map from cache: %s", urlKey)
		return outOrg, nil
	}

	db := GetSaasDBClient()
	so, err := organization.GetSaaSInfo(ctx, db)
	if err != nil {
		return nil, err
	}

	host := utils.CoalesceString(&so.Org.DBHost, &config.Config.Database.Host)
	orgDBCreds := config.DatabaseConfig{
		Host:     host,
		Database: so.Org.Database,
		Port:     config.Config.Database.Port,
		User:     config.Config.Database.User,
		Password: config.Config.Database.Password,
	}
	orgDBClient, err := postgres.GetDBFromConfig(ctx, orgDBCreds)
	if err != nil {
		log.Errorf("GetDBFromConfig: %s\n", err)
		return nil, err
	}

	orgResources := &OrgResources{
		Info: so,
		DB:   orgDBClient,
	}

	orgmap[urlKey] = orgResources

	return orgResources, nil
}

// GetDBClient return a configured DB client as a singleton
func GetSaasDBClient() *pgxpool.Pool {
	if saasDBClient != nil {
		return saasDBClient
	}

	lock.Lock()
	defer lock.Unlock()

	log.Debugf("Creating Master DBClient instance now.")

	// Connect to SurrealDB
	//"postgres://username:password@localhost:5432/database_name"
	host := fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
		config.Config.MasterDB.User,
		config.Config.MasterDB.Password,
		config.Config.MasterDB.Host,
		config.Config.MasterDB.Port,
		config.Config.MasterDB.Database)

	// host := fmt.Sprintf("ws://%s:%v/rpc", config.Config.Database.User, config.Config.Database.Password)
	db, err := postgres.GetDB(context.Background(), host)
	if err != nil {
		log.Error(err)
		return nil
	}

	saasDBClient = db

	return saasDBClient
}
