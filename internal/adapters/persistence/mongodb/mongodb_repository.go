package mongodb

import (
	"context"
	"log"
	"os"
	"strconv"
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
	var (
		mongodbAtlasPassword = os.Getenv("MONGODB_ATLAS_PASSWORD")
		mongodbName          = os.Getenv("MONGODB_NAME")
		mongodbPassword      = os.Getenv("MONGODB_PASSWORD")
		mongodbUser          = os.Getenv("MONGODB_USER")
		mongodbHost          = os.Getenv("MONGODB_HOST")
		mongodbPort          = os.Getenv("MONGODB_PORT")
		development          = os.Getenv("DEVELOPMENT")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	developmentBool, err := strconv.ParseBool(development)
	if err != nil {
		return nil, err
	}
	var client *mongo.Client
	if developmentBool {
		client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+mongodbUser+":"+mongodbPassword+"@"+mongodbHost+":"+mongodbPort+"/?authSource=admin"))
		if err != nil {
			return nil, err
		}
	} else {
		client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://eduardor2m:"+mongodbAtlasPassword+"@questao-certa.fvgdayd.mongodb.net/?retryWrites=true&w=majority"))
		if err != nil {
			return nil, err
		}
	}

	collection := client.Database(mongodbName)

	err = client.Ping(ctx, nil)

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
