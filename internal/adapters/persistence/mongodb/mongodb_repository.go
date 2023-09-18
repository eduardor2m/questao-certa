package mongodb

import (
	"context"
	"log"
	"os"
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
	dbName := os.Getenv("DB_NAME")
	dbAtlasPassword := os.Getenv("DB_ATLAS_PASSWORD")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbUser := os.Getenv("DB_USER")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPassword+"@"+dbHost+":"+dbPort))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://eduardor2m:"+dbAtlasPassword+"@questao-certa.fvgdayd.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName)

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
