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

	manifest := shared.LoadManifest()
	migrations.Apply(manifest)
}

// run the web server
func main() {
	ctx := config.NewContext()
	setup.SetupTestData(ctx)

	log.Infof("Starting GraphQL server listening on port %v", config.Config.Server.ServerPort)
	log.Fatal(appserv.Serve())
}
