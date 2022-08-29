package main

import (
	"context"
	"log"
	"time"

	"github.com/sinulingga23/go-pos/config"
	"github.com/sinulingga23/go-pos/implementation/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Printf("[DATABASE]: %s\b", err.Error())
	}

	client := database.Client()
	defer func() {
		if err := client.Disconnect(ctx);  err != nil {
			log.Printf("[DATABASE]: %s", err.Error())
		}
	}()

	cpn := repository.NewCategoryProductRepository(database)
	id, err := primitive.ObjectIDFromHex("630c65518b08ac4fd1bd6d83")
	if err != nil {
		log.Printf("[DATABASE]: %s", err.Error())
	}
	currentCategoryProduct, err := cpn.FindById(ctx, id)
	if err != nil {
		log.Printf("[DATABASE]: %s", err.Error())
	}
	log.Printf("[RESULT]: %v\n", currentCategoryProduct)
}