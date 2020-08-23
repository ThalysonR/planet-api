package main

import (
	"flag"
	"os"

	"github.com/thalysonr/planet-api/controllers"
	"github.com/thalysonr/planet-api/mappings"
	"github.com/thalysonr/planet-api/models"
)

// @title Planets API
// @version 1.0
// @description API para manter planetas
// @contact.name Thalyson
// @BasePath /api/v1
func main() {
	config := &controllers.AppConfig{}
	flag.StringVar(&config.ServerHost, "server-host", lookupEnvOrString("SERVER_HOST", config.ServerHost), "Endereço do host onde este serviço será executado")
	flag.StringVar(&config.MongoURI, "mongo-uri", lookupEnvOrString("MONGO_URI", config.MongoURI), "Endereço do mongoDB")
	flag.StringVar(&config.StarWarsAPI, "star-wars-api", lookupEnvOrString("SW_API", config.StarWarsAPI), "Endereço da API do Star Wars")

	repo, err := models.NewDB(config.MongoURI)
	if err != nil {
		panic(err)
	}
	defer repo.Disconnect()
	router := mappings.CreateRouter(repo, config)
	router.Run(":8080")
}

func lookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
