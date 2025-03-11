package main

import (
	"github.com/cogniia/core-api-template/src/config"
	server_service "github.com/cogniia/core-api-template/src/server/service"
)

// @title						Core API template
// @version					1.0
// @description				Template to use in your challenge
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	config.Load()
	// integration.Load()
	server_service.Serve()
}
