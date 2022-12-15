package main

import (
	"event-management/app"
	"event-management/config"
	"event-management/db"
	"event-management/docs"
	"event-management/utils"
	"event-management/utils/env"
	"flag"
	"fmt"
	"strings"
)

// @title                      Event Management
// @version                    1.0
// @host                       localhost
// @BasePath                   /
// @securityDefinitions.apikey ApiBearerToken
// @in                         header
// @name                       Authorization
func main() {
	env.Setup()

	db.Setup()
	defer db.Connection.Close()

	db.Migrate()

	cfg := config.Main{
		AppName: flag.String("app-name", env.Get("APP_NAME", "Event Management"), "Application Name"),
		Port:    flag.String("port", env.Get("APP_PORT", "3000"), "Port to listen on"),
		Prod:    flag.Bool("prod", utils.ParseBool(env.Get("IS_PRODUCTION", "false")), "Enable prefork in Production"),
	}

	flag.Parse()

	// Set dynamic host for Swagger documentation
	fullUrl := strings.Split(env.Get("APP_URL", "http://127.0.0.1:3000"), "://")
	scheme := fullUrl[0]
	host := fullUrl[1]
	if port := env.Get("APP_PORT", "3000"); port != "80" && host == "localhost" {
		host += fmt.Sprintf(":%s", port)
	}

	version := "v" + utils.GetCurrentVersion()

	if env.Get("IS_PRODUCTION", "false") != "true" {
		version += "-dev"
	}

	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Schemes = []string{scheme}
	docs.SwaggerInfo.Version = version

	app.Setup(cfg)
}
