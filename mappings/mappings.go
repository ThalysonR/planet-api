package mappings

import (
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

// CreateRouter creates the router with the endpoints the api exposes
func CreateRouter(repo models.IRepository, config *controllers.AppConfig) *gin.Engine {
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
		v1.GET("/planetas/:id", controller.BuscaPlanetaPorID)
		v1.GET("/planetas", controller.BuscaPlanetasPaginado)
		v1.POST("/planetas", controller.InserirPlaneta)
		v1.DELETE("/planetas/:id", controller.RemoverPlaneta)
	}

	return router
}
