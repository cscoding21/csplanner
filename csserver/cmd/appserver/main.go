package main

import (
	"csserver/internal/appserv"
	"csserver/internal/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	config.InitConfig()

	log.Infof("Starting GraphQL server listening on port %v", config.Config.Server.ServerPort)
	log.Fatal(appserv.Serve())
}
