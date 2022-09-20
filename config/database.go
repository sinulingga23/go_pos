package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongoDb(ctx context.Context) (*mongo.Database, error) {

	mongoDbUri := os.Getenv("MONGO_DB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUri))
	if err != nil {
		return nil, err
	}

	mongoDbName := os.Getenv("MONGO_DB_NAME")
	database := client.Database(mongoDbName)
	if err := database.Client().Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return database, nil
}