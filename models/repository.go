package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// IRepository contem todas as interfaces de repositórios
type IRepository interface {
	IPlanetaRepository
	Connect() error
	Disconnect() error
}

type repository struct {
	client                *mongo.Client
	connectionName        string
	databaseName          string
	planetaCollectionName string
}

// NewRepository cria uma instância de repositório e retorna sua interface
func NewRepository(connectionName string) (IRepository, error) {
	repository := &repository{
		client:                nil,
		connectionName:        connectionName,
		databaseName:          "planeta-api",
		planetaCollectionName: "planetas",
	}
	repository.Connect()
	return repository, nil
}

// Disconnect fecha a conexão com o banco de dados
func (repository *repository) Disconnect() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return repository.client.Disconnect(ctx)
}

// Connect abre a conexão com o banco de dados
func (repository *repository) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", repository.connectionName)))
	if err != nil {
		return err
	}
	repository.client = client
	return nil
}

func (repository *repository) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return repository.client.Ping(ctx, readpref.Primary())
}
