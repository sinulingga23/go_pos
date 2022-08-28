package main

import (
	"log"

	"github.com/sinulingga23/go-pos/config"
)

func main() {
	_, err := config.ConnectToMongoDb()
	if err != nil {
		log.Printf("[DATABASE]: %s\b", err.Error())
	}
}