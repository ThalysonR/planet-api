package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertResult é o resultado de uma inserção
type InsertResult struct {
	ID primitive.ObjectID
}

// NewDB cria uma instância de repositório e retorna sua interface
func NewDB(connectionName string) (IRepository, error) {
	repository := &repository{client: nil, connectionName: connectionName}
	repository.Connect()
	return repository, nil
}
