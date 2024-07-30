package main

import (
	"csserver/cmd/appserver/setup"
	"csserver/internal/appserv"
	"csserver/internal/config"
	"csserver/migrations"

	"github.com/cscoding21/csmig/shared"
	log "github.com/sirupsen/logrus"
)

// setup global configuration values
func init() {
	config.InitConfig()
	config.InitLogger()

	migConfig := GetMigrationConfig()
	migrations.Apply(migConfig)
}

// run the web server
func main() {
	ctx := config.NewContext()
	setup.SetupTestData(ctx)

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
