package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/thalysonr/planet-api/httpclient"
	"github.com/thalysonr/planet-api/models"
)

// Controller é a base de todos os controllers
type Controller struct {
	planetasRepository models.IPlanetaRepository
	httpClient         httpclient.IHTTPClient
	config             *AppConfig
	validate           *validator.Validate
}

// NewController retorna uma instancia de Controller
func NewController(repository models.IRepository, client httpclient.IHTTPClient, config *AppConfig) *Controller {
	return &Controller{
		planetasRepository: repository,
		config:             config,
		httpClient:         client,
		validate:           validator.New(),
	}
}

// AppConfig contem as configurações necessárias ao web server
type AppConfig struct {
	StarWarsAPI string
	ServerHost  string
	MongoURI    string
}
