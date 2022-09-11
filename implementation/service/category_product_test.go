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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func TestCategoryProductService_FindById_Exists(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	_, err = database.Collection(repository.CategoryProductCollection).DeleteMany(ctx, struct{}{})
	if err != nil {
		log.Fatal(err.Error())
	}
	
	wantCategoryName := "Pakaian"
	wantDescription := "Menjual berbagai pakaian yang paling update."

	categoryProductRepository := repository.NewCategoryProductRepository(database)
	categoryProductService := NewCategoryProductService(categoryProductRepository)

	createdCategoryProduct, err := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Pakaian",
		Description: "Menjual berbagai pakaian yang paling update.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantId := createdCategoryProduct.Id

	currentCategoryProduct, err := categoryProductService.FindById(createdCategoryProduct.Id)
	if err != nil {
		log.Fatal(err.Error())
	}

	if strings.Compare(wantId, currentCategoryProduct.Id) != 0 {
		log.Fatalf("got %q want %q\n", currentCategoryProduct.Id, wantId)
	}

	if strings.Compare(wantCategoryName, currentCategoryProduct.CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", currentCategoryProduct.CategoryName, wantCategoryName)
	}

	if strings.Compare(wantDescription, currentCategoryProduct.Description) != 0 {
		log.Fatalf("got %q want %q\n", currentCategoryProduct.Description, wantDescription)
	}
}

func TestCategoryProductService_FindById_NotExists(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(repository.CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantError1 := definition.ErrDataNotFound
	wantError2 := definition.ErrDataNotFound

	categoryProductRepository := repository.NewCategoryProductRepository(database)
	categoryProductService := NewCategoryProductService(categoryProductRepository)

	_, err1 := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Pinjaman",
		Description: "Butuhan membayar pinjaman atau mengajukan angsuran ? Di sini tempatnya.",
	})
	if err1 != nil {
		log.Fatal(err1.Error())
	}

	_, err2 := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Kesehatan",
		Description: "Menyediakan berbagai macam jenis obat.",
	})
	if err2 != nil {
		log.Fatal(err2.Error())
	}

	_, err1 = categoryProductService.FindById(primitive.NewObjectID().Hex())
	_, err2 = categoryProductService.FindById(primitive.NewObjectID().Hex())

	if wantError1 != err1 {
		log.Fatalf("got %q want %q\n", err1.Error(), wantError1.Error())
	}

	if wantError2 != err2 {
		log.Fatalf("got %q want %q\n", err2.Error(), wantError2.Error())
	}
}