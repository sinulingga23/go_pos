package main

import (
	"context"
	"log"
	"time"

	"github.com/sinulingga23/go-pos/config"
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
}