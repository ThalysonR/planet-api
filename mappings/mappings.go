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
	_ "github.com/thalysonr/planet-api/docs" //swagger docs
)

// CreateRouter creates the router with the endpoints the api exposes
func CreateRouter(serverHost string) *gin.Engine {
	router := gin.New()
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

	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", serverHost))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(gin.Recovery())
	router.Use(logger.SetLogger())
	router.Use(cors.New(corsConfig))
	pprof.Register(router)
	return router
}
