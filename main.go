package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sinulingga23/go-pos/config"
	"github.com/sinulingga23/go-pos/handler"
	"github.com/sinulingga23/go-pos/implementation/repository"
	"github.com/sinulingga23/go-pos/implementation/service"
	thirdparty_service "github.com/sinulingga23/go-pos/implementation/service/third-party"
	"google.golang.org/api/option"
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

	clientStorage, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("STORAGE_CREDENTIAL_FILE")))
	if err != nil {
		log.Printf("[CLIENT-STORAGE]: %s", err.Error())
	}

	categoryProductRepository := repository.NewCategoryProductRepository(database)
	productRepository := repository.NewProductRepository(database)

	categoryProductService := service.NewCategoryProductService(categoryProductRepository)
	cloudStorageService := thirdparty_service.NewCloudStorageService(clientStorage)

	productService := service.NewProductService(productRepository, categoryProductRepository, cloudStorageService)
	
	handler := handler.NewHandler(categoryProductService, productService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/api/v1/category-product", handler.CreateCategoryProduct)
	router.Get("/api/v1/category-product/{id}", handler.FindCategoryById)
	router.Put("/api/v1/category-product/{id}", handler.UpdateCategoryProductById)
	router.Delete("/api/v1/category-product/{id}", handler.DeleteCategoryProductById)
	router.Get("/api/v1/category-product", handler.GetAllCategoryProduct)
	
	router.Post("/api/v1/product", handler.CreateProduct)
	router.Post("/api/v1/product/{id}/add-images", handler.AddImagesToProduct)

	log.Print("Running on :8081\n")

	if err := http.ListenAndServe("0.0.0.0:8081", router); err != nil {
		log.Fatalf("[ERROR:] %s\v", err)
	}
}