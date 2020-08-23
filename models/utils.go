package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type count struct {
	Count int `bson:"count"`
}

type PaginaResultado struct {
	Anterior     *string     `json:"anterior"`
	Dados        interface{} `json:"dados"`
	Pagina       int         `json:"pagina"`
	Proxima      *string     `json:"proxima"`
	TotalPaginas int         `json:"total_paginas"`
}

// ResultadoInsert é o resultado de uma inserção
type ResultadoInsert struct {
	ID primitive.ObjectID
}

func agregacaoPaginada(fieldName *string, fieldValue *string, skip int, limit int) bson.D {
	var match bson.D
	if fieldName == nil || fieldValue == nil {
		match = bson.D{{}}
	} else {
		match = bson.D{{*fieldName, bson.D{{"$regex", *fieldValue}, {"$options", "i"}}}}
	}
	data := bson.A{bson.D{{"$match", match}}, bson.D{{"$skip", skip}}, bson.D{{"$limit", limit}}}
	count := bson.A{bson.D{{"$count", "count"}}}
	facet := bson.D{{"$facet", bson.D{{"data", data}, {"count", count}}}}
	return facet
}
