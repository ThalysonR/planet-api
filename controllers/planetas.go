package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/thalysonr/planet-api/models"
)

type swResponse struct {
	Count   int                `json:"count"`
	Results []swResponseResult `json:"results"`
}

type swResponseResult struct {
	Name  string   `json:"name"`
	Films []string `json:"films"`
}

// BuscaPlanetaPorID é o controller que faz busca de um unico planeta pelo ID
// @Summary Busca planeta
// @Description Busca um planeta por ID
// @Tags Planetas
// @Accept json
// @Produce json
// @Param id path string true "id do planeta a ser buscado"
// @Success 200 {object} models.Planeta
// @Failure 404 {string} string "Not found"
// @Router /planetas/{id} [get]
func (ct *Controller) BuscaPlanetaPorID(c *gin.Context) {
	id := c.Param("id")
	planeta, err := ct.planetasRepository.BuscaPlanetaPorID(id)
	if err != nil {
		log.Debug().Err(err).Send()
		c.JSON(http.StatusNotFound, &gin.H{"erro": "id não encontrado"})
		return
	}
	c.JSON(http.StatusOK, planeta)
}

// BuscaPlanetasPaginado é o controller que faz busca de planetas
// @Summary Busca planetas
// @Description Busca paginada de planetas
// @Tags Planetas
// @Accept json
// @Produce json
// @Param busca query string false "buscar por nome do planeta"
// @Param skip query int false "pular X itens no resultado"
// @Param limit query int false "limite de itens por pagina"
// @Success 200 {object} models.PaginaResultado
// @Failure 500 {string} string "Internal server error"
// @Router /planetas [get]
func (ct *Controller) BuscaPlanetasPaginado(c *gin.Context) {
	var filtroCampoNome *string
	var filtroCampoValor *string
	var pCampoNome string
	iLimit := 5
	iSkip := 0
	pNome := c.Query("busca")
	pSkip := c.Query("skip")
	pLimit := c.Query("limit")
	if pNome != "" {
		pCampoNome = "nome"
		filtroCampoValor = &pNome
	}
	filtroCampoNome = &pCampoNome
	if tmplimit, err := strconv.ParseInt(pLimit, 10, 32); err == nil {
		iLimit = int(math.Max(float64(tmplimit), 1))
	}
	if tmpskip, err := strconv.ParseInt(pSkip, 10, 32); err == nil {
		iSkip = int(math.Max(float64(tmpskip), 0))
	}
	planetas, err := ct.planetasRepository.BuscaPlanetasPaginado(filtroCampoNome, filtroCampoValor, iSkip, iLimit)
	if err != nil {
		log.Debug().Err(err).Msg("falha ao buscar planetas")
		c.JSON(http.StatusInternalServerError, &gin.H{"erro": "falha ao buscar planetas"})
		return
	}
	prox := fmt.Sprintf("%s/api/v1/%s?busca=%s&skip=%d&limit=%d", ct.config.ServerHost, ct.controllerNames.PlanetasController, pNome, iSkip+iLimit, iLimit)
	anterior := fmt.Sprintf("%s/api/v1/%s?busca=%s&skip=%d&limit=%d", ct.config.ServerHost, ct.controllerNames.PlanetasController, pNome, iSkip-iLimit, iLimit)
	if planetas.Pagina < planetas.TotalPaginas {
		planetas.Proxima = &prox
	}
	if planetas.Pagina > 1 {
		planetas.Anterior = &anterior
	}
	c.JSON(http.StatusOK, planetas)
}

// InserirPlaneta é o controller que insere planetas
// @Summary Inserir Planeta
// @Description Insere um novo planeta
// @Tags Planetas
// @Accept json
// @Produce json
// @Param planeta_input body models.PlanetaInput true "Dados de um planeta"
// @Success 201 {object} models.ResultadoInsert
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal server error"
// @Router /planetas [post]
func (ct *Controller) InserirPlaneta(c *gin.Context) {
	var input models.PlanetaInput
	err := c.BindJSON(&input)
	if err != nil {
		log.Debug().Err(err).Send()
		c.JSON(http.StatusBadRequest, &gin.H{"erro": "falha ao decodificar entrada"})
		return
	}
	err = ct.validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"erro": "falha na validação dos campos de entrada"})
		return
	}
	res, err := ct.httpClient.HTTPRequest("GET", ct.config.StarWarsAPI, fmt.Sprintf("planets?search=%s", input.Nome), nil)
	if err == nil {
		var swResp swResponse
		err = json.Unmarshal(res, &swResp)
		if err == nil {
			if len(swResp.Results) > 0 && strings.ToLower(input.Nome) == strings.ToLower(swResp.Results[0].Name) {
				input.AparicoesEmFilmes = len(swResp.Results[0].Films)
			}
		}
	}
	insertRes, err := ct.planetasRepository.InserirPlaneta(input)
	if err != nil {
		log.Debug().Err(err).Send()
		c.JSON(http.StatusInternalServerError, &gin.H{"erro": "falha ao inserir planeta"})
		return
	}
	c.JSON(http.StatusCreated, insertRes)
}

// RemoverPlaneta deleta um planeta do banco usando seu ID
// @Summary Remover Planeta
// @Description Remove um planeta usando seu ID
// @Tags Planetas
// @Accept json
// @Produce json
// @Param id path string true "ID do planeta a ser deletado"
// @Success 200 {string} string "Item removido"
// @Failure 400 {string} string "Bad request"
// @Router /planetas/{id} [delete]
func (ct *Controller) RemoverPlaneta(c *gin.Context) {
	id := c.Param("id")
	err := ct.planetasRepository.RemoverPlaneta(id)
	if err != nil {
		log.Debug().Err(err).Send()
		c.JSON(http.StatusBadRequest, &gin.H{"erro": "planeta não encontrado"})
		return
	}
	c.JSON(http.StatusOK, &gin.H{"msg": "Item deletado"})
}
