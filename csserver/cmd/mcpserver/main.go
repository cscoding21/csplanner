package main

import (
	"csserver/internal/appserv/csmap"
	"csserver/internal/appserv/csmcp"
	"csserver/internal/config"
	"fmt"

	"github.com/mark3labs/mcp-go/server"

	log "github.com/sirupsen/logrus"
)

// setup global configuration values
func init() {
	config.InitConfig()
	config.InitLogger()

	csmap.ClearAllCaches()
}

// run the MCP server
func main() {
	port := 6000

	log.Infof("Starting MCP server listening on port %v", port)

	mcpServer := csmcp.GetCSMCPServer()
	httpServer := server.NewStreamableHTTPServer(mcpServer)
	log.Infof("MCP HTTP server listening on :%v/mcp", port)
	if err := httpServer.Start(fmt.Sprintf(":%v", port)); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
