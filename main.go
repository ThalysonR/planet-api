package main

import (
	"flag"
	"os"

	"github.com/thalysonr/planet-api/mappings"
)

var (
	serverHost string
)

// @title Planets API
// @version 1.0
// @description API to store and retrieve planets information
// @contact.name Thalyson
// @BasePath /api/v1
func main() {
	flag.StringVar(&serverHost, "server-host", lookupEnvOrString("SERVER_HOST", serverHost), "Address of the host where this service will be run")

	router := mappings.CreateRouter(serverHost)
	router.Run(":8080")
}

func lookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
