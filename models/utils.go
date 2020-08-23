package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type count struct {
	Count int `bson:"count"`
}

type PaginaResultado struct {
	Dados        interface{} `json:"dados"`
	Pagina       int         `json:"pagina"`
	TotalPaginas int         `json:"total_paginas"`
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
