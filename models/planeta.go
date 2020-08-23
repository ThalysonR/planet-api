package models

import (
	"context"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// IPlanetaRepository contem as possíveis operações no model planeta
type IPlanetaRepository interface {
	BuscaPlanetaPorID(ID string) (*Planeta, error)
	BuscaPlanetasPaginado(campoNome *string, campoValor *string, skip int, limit int) (*PaginaResultado, error)
	InserirPlaneta(input PlanetaInput) (*InsertResult, error)
	RemoverPlaneta(ID string) error
}

// Planeta representa um planeta
type Planeta struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	PlanetaInput `bson:",inline"`
	CriadoEm     time.Time `bson:"criado_em" json:"criado_em"`
	AtualizadoEm time.Time `bson:"atualizado_em" json:"atualizado_em"`
}

// PlanetaInput é o input esperado pela API para inserir um planeta
type PlanetaInput struct {
	AparicoesEmFilmes int    `bson:"aparicoes_em_filmes" json:"aparicoes_em_filmes"`
	Clima             string `bson:"clima" json:"clima" validate:"required"`
	Nome              string `bson:"nome" json:"nome" validate:"required"`
	Terreno           string `bson:"terreno" json:"terreno" validate:"required"`
}

type respostaPlanetaAgregacao struct {
	Dados []*Planeta `bson:"data"`
	Count []*count   `bson:"count"`
}

var (
	databaseName   = "planeta-api"
	collectionName = "planetas"
)

func (repo *repository) BuscaPlanetaPorID(ID string) (*Planeta, error) {
	err := repo.Ping()
	if err != nil {
		repo.Connect()
		return nil, err
	}
	collection := repo.client.Database(databaseName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var planeta *Planeta
	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", oID}}
	err = collection.FindOne(ctx, filter).Decode(&planeta)
	if err != nil {
		return nil, err
	}
	return planeta, nil
}

func (repo *repository) BuscaPlanetasPaginado(campoNome *string, campoValor *string, skip int, limit int) (*PaginaResultado, error) {
	err := repo.Ping()
	if err != nil {
		repo.Connect()
		return nil, err
	}
	collection := repo.client.Database(databaseName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	agreg := agregacaoPaginada(campoNome, campoValor, skip, limit)
	cur, err := collection.Aggregate(ctx, mongo.Pipeline{agreg})
	if err != nil {
		return nil, err
	}
	var resposta respostaPlanetaAgregacao
	cur.Next(ctx)
	err = cur.Decode(&resposta)
	if err != nil {
		return nil, err
	}
	cur.Close(ctx)
	resultados := &PaginaResultado{
		Dados:        resposta.Dados,
		Pagina:       skip / limit,
		TotalPaginas: int(math.Ceil(float64(resposta.Count[0].Count) / float64(limit))),
	}
	return resultados, nil
}

func (repo *repository) InserirPlaneta(input PlanetaInput) (*InsertResult, error) {
	err := repo.Ping()
	if err != nil {
		repo.Connect()
		return nil, err
	}
	collection := repo.client.Database(databaseName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	id := primitive.NewObjectID()
	planeta := &Planeta{
		PlanetaInput: input,
		ID:           id,
		CriadoEm:     time.Now(),
		AtualizadoEm: time.Now(),
	}
	_, err = collection.InsertOne(ctx, planeta)
	if err != nil {
		return nil, err
	}
	return &InsertResult{ID: id}, nil
}

func (repo *repository) RemoverPlaneta(ID string) error {
	err := repo.Ping()
	if err != nil {
		repo.Connect()
		return err
	}
	collection := repo.client.Database(databaseName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", oID}}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
