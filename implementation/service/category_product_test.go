package service

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/sinulingga23/go-pos/config"
	"github.com/sinulingga23/go-pos/definition"
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

func TestCategoryProductService_Create_TheFieldsIsEmpty(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(repository.CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantError1 := definition.ErrBadRequest
	wantError2 := definition.ErrBadRequest

	categoryProductRepository := repository.NewCategoryProductRepository(database)
	categoryProductService := NewCategoryProductService(categoryProductRepository)

	_, err1 :=  categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "",
		Description: "",
	})

	_, err2 := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "   ",
		Description: "       ",
	})

	if wantError1 != err1 {
		log.Fatalf("got %q want %q\n", err1.Error(), wantError1.Error())
	}

	if wantError2 != err2 {
		log.Fatalf("got %q want %q\n", err2.Error(), wantError2.Error())
	}
}

func TestCategoryProductService_Create_IsThereFieldsEmptyAndNot(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(repository.CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantError1 := definition.ErrBadRequest
	wantError2 := definition.ErrBadRequest

	categoryProductRepository := repository.NewCategoryProductRepository(database)
	categoryProductService := NewCategoryProductService(categoryProductRepository)

	_, err1 :=  categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "",
		Description: "Menyediakan semua kebutuhan hiburanmu!.",
	})

	_, err2 := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Hiburan",
		Description: "",
	})

	if wantError1 != err1 {
		log.Fatalf("got %q want %q\n", err1.Error(), wantError1.Error())
	}

	if wantError2 != err2 {
		log.Fatalf("got %q want %q\n", err2.Error(), wantError2.Error())
	}
}