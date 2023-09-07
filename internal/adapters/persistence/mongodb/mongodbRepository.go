package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type connectorManager interface {
	getConnection() (*mongo.Database, error)
	closeConnection(conn *mongo.Database)
}

var _ connectorManager = (*DatabaseConnectorManager)(nil)

type DatabaseConnectorManager struct{}

func (dcm *DatabaseConnectorManager) getConnection() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		return nil, err
	}

	collection := client.Database("questao-certa")

	if err != nil {
		log.Fatal(err)
	}

	return collection, nil
}

func (dcm *DatabaseConnectorManager) closeConnection(conn *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	conn.Client().Disconnect(ctx)
}
