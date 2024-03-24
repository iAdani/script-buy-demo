package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

// Create Singleton Class for Database Connection
type databaseConnection struct {
	Client *mongo.Client
	Ctx    context.Context
}

var db *databaseConnection

func Connect() *databaseConnection {
	atlasUri := os.Getenv("ATLAS_URI")
	if db == nil {
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().ApplyURI(atlasUri).SetServerAPIOptions(serverAPIOptions)
		connCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		client, err := mongo.Connect(connCtx, clientOptions)
		if err != nil {
			panic(err)
		}
		db = &databaseConnection{Client: client}
	}
	return db
}

func CloseDatabaseConnection() {
	if db != nil {
		err := db.Client.Disconnect(db.Ctx)
		if err != nil {
			panic(err)
		}
	}
}
