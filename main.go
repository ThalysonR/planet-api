package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thalysonr/planet-api/controllers"
	_ "github.com/thalysonr/planet-api/docs" //swagger docs
	"github.com/thalysonr/planet-api/httpclient"
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

	repo, err := models.NewRepository(config.MongoURI)
	if err != nil {
		panic(err)
	}
	defer repo.Disconnect()
	router := createRouter(repo, config)
	router.Run(":8080")
}

func lookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func createRouter(repo models.IRepository, config *controllers.AppConfig) *gin.Engine {
	router := gin.New()
	controller := controllers.NewController(repo, httpclient.NewBreaker(), config)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", config.ServerHost))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(gin.Recovery())
	router.Use(logger.SetLogger())
	router.Use(cors.New(corsConfig))
	pprof.Register(router)

	api := router.Group("api")
	v1 := api.Group("v1")
	{
		v1.GET("/", controller.Raiz)
		v1.GET("/planetas/:id", controller.BuscaPlanetaPorID)
		v1.GET("/planetas", controller.BuscaPlanetasPaginado)
		v1.POST("/planetas", controller.InserirPlaneta)
		v1.DELETE("/planetas/:id", controller.RemoverPlaneta)
	}

	return router
}
