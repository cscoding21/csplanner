package main

import (
	"csserver/internal/appserv"
	"csserver/internal/appserv/csmap"
	"csserver/internal/config"

	"github.com/cscoding21/csmig/shared"
	log "github.com/sirupsen/logrus"
)

// setup global configuration values
func init() {
	config.InitConfig()
	config.InitLogger()

	csmap.ClearAllCaches()

	//migConfig := GetMigrationConfig()
	//migrations.Apply(migConfig)
}

// run the web server
func main() {
	log.Infof("Starting GraphQL server listening on port %v", config.Config.Server.ServerPort)
	log.Fatal(appserv.Serve())
}

func GetMigrationConfig() shared.MigratorConfig {
	return shared.MigratorConfig{
		DatabaseStrategyName: "surrealdb",
		GeneratorPath:        "migrations",
		GeneratorPackage:     "migrations",
		ImplementationName:   "csplanner",
		DBConfig: shared.DatabaseConfig{
			Name:      "surrealdb",
			Host:      config.Config.Database.Host,
			Port:      config.Config.Database.Port,
			User:      config.Config.Database.User,
			Password:  config.Config.Database.Password,
			Database:  config.Config.Database.Database,
			Namespace: config.Config.Database.Namespace,
		},
	}
}
