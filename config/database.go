package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongoDb(ctx context.Context) (*mongo.Database, error) {
	mongoDbUser := os.Getenv("MONGO_DB_USER")
	mongoDbPassword := os.Getenv("MONGO_DB_PASSWORD")
	mongoDbCluster := os.Getenv("MONGO_DB_CLUSTER")
	mongoDbName := os.Getenv("MONGO_DB_NAME")
	
	mongoUri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", mongoDbUser, mongoDbPassword, mongoDbCluster)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	database := client.Database(mongoDbName)
	if err := database.Client().Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return database, nil
}