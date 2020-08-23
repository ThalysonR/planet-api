package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/thalysonr/planet-api/httpclient"
	"github.com/thalysonr/planet-api/models"
)

func TestBuscaPlanetaPorID(t *testing.T) {
	Convey("Dado um ID", t, func() {
		id := "id_certo"
		planetaMock := &models.PlanetaMock{
			BuscaPlanetaPorIDMock: func(ID string) (*models.Planeta, error) {
				if ID == id {
					return nil, nil
				}
				return nil, errors.New("ID errado")
			},
		}
		controller := &Controller{
			planetasRepository: planetaMock,
		}
		Convey("Quando o repositório encontra o planeta com este ID, a função deve retornar 200", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = append(c.Params, gin.Param{
				Key:   "id",
				Value: "id_certo",
			})
			controller.BuscaPlanetaPorID(c)
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Quando o repositório não encontra o planeta com este ID, a função deve retornar um código diferente de 200", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = append(c.Params, gin.Param{
				Key:   "id",
				Value: "id_errado",
			})
			controller.BuscaPlanetaPorID(c)
			So(w.Code, ShouldNotEqual, 200)
		})
	})
}

func TestBuscaPlanetaPaginado(t *testing.T) {
	Convey("Dado um repositório de planetas", t, func() {
		planetaMock := &models.PlanetaMock{
			BuscaPlanetasPaginadoMock: func(campoNome *string, campoValor *string, skip int, limit int) (*models.PaginaResultado, error) {
				if *campoValor == "erro" {
					return nil, errors.New("Erro")
				}
				return nil, nil
			},
		}
		controller := &Controller{
			planetasRepository: planetaMock,
		}
		Convey("Caso o repositório retorne com itens, a função deve retornar 200", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "http://example.com?busca=teste&skip=0&limit=5", nil)
			c, _ := gin.CreateTestContext(w)
			c.Request = r
			controller.BuscaPlanetasPaginado(c)
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Caso o repositorio retorne um erro, a função deve retornar um código diferente de 200", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "http://example.com?busca=erro", nil)
			c, _ := gin.CreateTestContext(w)
			c.Request = r
			controller.BuscaPlanetasPaginado(c)
			So(w.Code, ShouldNotEqual, 200)
		})
	})
}

func TestInserePlaneta(t *testing.T) {
	Convey("Dado uma entrada", t, func() {
		httpclientMock := &httpclient.HTTPClientMock{
			HTTPRequestMock: func(method string, url string, path string, body io.Reader) ([]byte, error) {
				resp := &swResponse{
					Results: []swResponseResult{
						{
							Name: "Certo",
							Films: []string{
								"teste",
								"teste",
								"teste",
							},
						},
					},
				}
				bytes, _ := json.Marshal(resp)
				return bytes, nil
			},
		}
		planetaRepoMock := &models.PlanetaMock{
			InserirPlanetaMock: func(planeta models.PlanetaInput) (*models.InsertResult, error) {
				if planeta.Nome == "erro" {
					return nil, errors.New("Erro")
				}
				return nil, nil
			},
		}
		appConfig := &AppConfig{
			StarWarsAPI: "teste",
		}
		Convey("Caso não seja um JSON válido, a função deve retornar um código diferente de 201", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			controller := &Controller{}
			controller.InserirPlaneta(c)
			So(w.Code, ShouldNotEqual, 201)
		})
		Convey("Caso faltem campos necessários, a função deve retornar um código diferente de 201", func() {
			input := &models.PlanetaInput{}
			inputBytes, _ := json.Marshal(input)
			inputReader := bytes.NewReader(inputBytes)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://example.com", inputReader)
			c, _ := gin.CreateTestContext(w)
			c.Request = r
			controller := &Controller{validate: validator.New()}
			controller.InserirPlaneta(c)
			So(w.Code, ShouldNotEqual, 201)
		})
		Convey("Caso o repositório retorne erro, a função deve retornar um código diferente de 201", func() {
			input := &models.PlanetaInput{
				Clima:   "Árido",
				Nome:    "erro",
				Terreno: "Deserto",
			}
			inputBytes, _ := json.Marshal(input)
			inputReader := bytes.NewReader(inputBytes)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://example.com", inputReader)
			c, _ := gin.CreateTestContext(w)
			c.Request = r
			httpclientMock.HTTPRequest("GET", "teste", "teste", nil)
			controller := &Controller{validate: validator.New(), httpClient: httpclientMock, planetasRepository: planetaRepoMock, config: appConfig}
			controller.InserirPlaneta(c)
			fmt.Println(w.Code)
			So(w.Code, ShouldNotEqual, 201)
		})
		Convey("Caso o repositório não retorne erro, a função deve retornar 201", func() {
			input := &models.PlanetaInput{
				Clima:   "Árido",
				Nome:    "Certo",
				Terreno: "Deserto",
			}
			inputBytes, _ := json.Marshal(input)
			inputReader := bytes.NewReader(inputBytes)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://example.com", inputReader)
			c, _ := gin.CreateTestContext(w)
			c.Request = r
			httpclientMock.HTTPRequest("GET", "teste", "teste", nil)
			controller := &Controller{validate: validator.New(), httpClient: httpclientMock, planetasRepository: planetaRepoMock, config: appConfig}
			controller.InserirPlaneta(c)
			fmt.Println(w.Code)
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestRemoverPlaneta(t *testing.T) {
	Convey("Dado um id", t, func() {
		id := "id_certo"
		planetaMock := &models.PlanetaMock{
			RemoverPlanetaMock: func(ID string) error {
				if ID == id {
					return nil
				}
				return errors.New("ID errado")
			},
		}
		controller := &Controller{
			planetasRepository: planetaMock,
		}
		Convey("Caso o repositório retorne erro, a funçao deve retornar um código diferente de 200", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = append(c.Params, gin.Param{
				Key:   "id",
				Value: "id_errado",
			})
			controller.RemoverPlaneta(c)
			So(w.Code, ShouldNotEqual, 200)
		})
		Convey("Caso o repositório não retorne erro, a função deve retornar código 200", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = append(c.Params, gin.Param{
				Key:   "id",
				Value: "id_certo",
			})
			controller.RemoverPlaneta(c)
			So(w.Code, ShouldEqual, 200)
		})
	})
}
