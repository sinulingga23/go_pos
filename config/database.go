package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongoDb() (*mongo.Database, error) {
	mongoDbUser := os.Getenv("MONGO_DB_USER")
	mongoDbPassword := os.Getenv("MONGO_DB_PASSWORD")
	mongoDbHost := os.Getenv("MONGO_DB_HOST")
	mongoDbPort := os.Getenv("MONGO_DB_PORT")
	mongoDBName := os.Getenv("MONGO_DB_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoDbUser, mongoDbPassword, mongoDbHost, mongoDbPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	database := client.Database(mongoDBName)
	if err := database.Client().Ping(ctx, readpref.Primary()); err != nil {
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	log.Printf("[DATABASE]: Connect successfully to MongoDB!\n")
	return database, nil
}