// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Thalyson"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/planetas": {
            "get": {
                "description": "Busca paginada de planetas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Planetas"
                ],
                "summary": "Busca planetas",
                "parameters": [
                    {
                        "type": "string",
                        "description": "buscar por nome do planeta",
                        "name": "busca",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "pular X itens no resultado",
                        "name": "skip",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limite de itens por pagina",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PaginaResultado"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Insere um novo planeta",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Planetas"
                ],
                "summary": "Inserir Planeta",
                "parameters": [
                    {
                        "description": "Dados de um planeta",
                        "name": "planeta_input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PlanetaInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.InsertResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/planetas/{id}": {
            "get": {
                "description": "Busca um planeta por ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Planetas"
                ],
                "summary": "Busca planeta",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id do planeta a ser buscado",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Planeta"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove um planeta usando seu ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Planetas"
                ],
                "summary": "Remover Planeta",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID do planeta a ser deletado",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Item removido",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.InsertResult": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "models.PaginaResultado": {
            "type": "object",
            "properties": {
                "dados": {
                    "type": "object"
                },
                "pagina": {
                    "type": "integer"
                },
                "total_paginas": {
                    "type": "integer"
                }
            }
        },
        "models.Planeta": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "atualizado_em": {
                    "type": "string"
                },
                "clima": {
                    "type": "string"
                },
                "criado_em": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "terreno": {
                    "type": "string"
                }
            }
        },
        "models.PlanetaInput": {
            "type": "object",
            "properties": {
                "clima": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "terreno": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Planets API",
	Description: "API para manter planetas",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
