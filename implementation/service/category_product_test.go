package service

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/sinulingga23/go-pos/config"
	"github.com/sinulingga23/go-pos/implementation/repository"
	"github.com/sinulingga23/go-pos/payload"
)


func TestCategoryProductService_Create_Success(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(repository.CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantCategoryName := "Kesehatan"
	wantDescription := "Menyediakan berbagai kebutuhan obat."

	categoryProductRepository := repository.NewCategoryProductRepository(database)
	categoryProductService := NewCategoryProductService(categoryProductRepository)

	createdCategoryProduct, err := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Kesehatan",
		Description: "Menyediakan berbagai kebutuhan obat.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	if strings.Compare(wantCategoryName, createdCategoryProduct.CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", createdCategoryProduct.CategoryName, wantCategoryName)
	}

	if strings.Compare(wantDescription, createdCategoryProduct.Description) != 0 {
		log.Fatalf("got %q want %q\n", createdCategoryProduct.Description, wantDescription)
	}
}