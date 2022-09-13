package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sinulingga23/go-pos/config"
	"github.com/sinulingga23/go-pos/handler"
	"github.com/sinulingga23/go-pos/implementation/repository"
	"github.com/sinulingga23/go-pos/implementation/service"
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

	categoryProductRepository := repository.NewCategoryProductRepository(database)
	categoryProductService := service.NewCategoryProductService(categoryProductRepository)
	
	handler := handler.NewHandler(categoryProductService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/api/v1/category-product", handler.CreateCategoryProduct)

	log.Printf("Running on :8085")

	if err := http.ListenAndServe(":8085", router); err != nil {
		log.Fatalf("[ERROR:] %s\v", err)
	}
}