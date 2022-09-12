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

func TestCategoryProductService_UpdateById_Success(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(repository.CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantCategoryName1 := "Kesehatan"
	wantDescription1 := "Menyediakan berbagai macam kebutuhan obat."

	wantCategoryName2 := "Elektronik"
	wantDescription2 := "Menyediakan berbagai macam kebutuhan elekronik, mulai dari laptop, speaker, sampai smartphone."

	wantCategoryName3 := "Pinjaman"
	wantDescription3 := "Menyediakan pembayaran instan sesuai kebutuhanmu."

	categoryProductRepository := repository.NewCategoryProductRepository(database)
	categoryProductService := NewCategoryProductService(categoryProductRepository)

	createdCategoryProduct1, err := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Ksehatan",
		Description: "Menyediakan brbgagai macam keubuthan obat.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	
	createdCategoryProduct2, err := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Elektroniksa",
		Description: "Menyediakan brbgagai.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	createdCategoryProduct3, err := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Pinjamanm",
		Description: "Menyediakan pembayaran instan sesuai kebutuhanmu...",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantId1 := createdCategoryProduct1.Id

	updatedCategoryProduct1, err := categoryProductService.UpdateById(createdCategoryProduct1.Id, payload.UpdateCategoryProductRequest{
		Id: createdCategoryProduct1.Id,
		CategoryName: "Kesehatan",
		Description: "Menyediakan berbagai macam kebutuhan obat.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	if strings.Compare(wantId1, updatedCategoryProduct1.Id) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct1.Id, wantId1)
	}

	if strings.Compare(wantCategoryName1, updatedCategoryProduct1.CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct1.CategoryName, wantCategoryName1)
	}

	if strings.Compare(wantDescription1, updatedCategoryProduct1.Description) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct1.Description, wantDescription1)
	}

	wantId2 := createdCategoryProduct2.Id

	updatedCategoryProduct2, err := categoryProductService.UpdateById(createdCategoryProduct2.Id, payload.UpdateCategoryProductRequest{
		Id: createdCategoryProduct2.Id,
		CategoryName: "Elektronik",
		Description: "Menyediakan berbagai macam kebutuhan elekronik, mulai dari laptop, speaker, sampai smartphone.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	if strings.Compare(wantId2, updatedCategoryProduct2.Id) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct2.Id, wantId2)
	}

	if strings.Compare(wantCategoryName2, updatedCategoryProduct2.CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct2.CategoryName, wantCategoryName2)
	}

	if strings.Compare(wantDescription2, updatedCategoryProduct2.Description) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct2.Description, wantDescription2)
	}

	wantId3 := createdCategoryProduct3.Id

	updatedCategoryProduct3, err := categoryProductService.UpdateById(createdCategoryProduct3.Id, payload.UpdateCategoryProductRequest{
		Id: createdCategoryProduct3.Id,
		CategoryName: "Pinjaman",
		Description: "Menyediakan pembayaran instan sesuai kebutuhanmu.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	
	if strings.Compare(wantId3, updatedCategoryProduct3.Id) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct3.Id, wantId3)
	}

	if strings.Compare(wantCategoryName3, updatedCategoryProduct3.CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct3.CategoryName, wantCategoryName3)
	}

	if strings.Compare(wantDescription3, updatedCategoryProduct3.Description) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct3.Description, wantDescription3)
	}
}

func TestCategoryProductService_DeleteById(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(repository.CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	categoryProductRepository := repository.NewCategoryProductRepository(database)
	categoryProductService := NewCategoryProductService(categoryProductRepository)

	createdCategoryProduct1, err := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Kesehatan",
		Description: "Menyediakan berbagai macam kebutuhan obat - obatan.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	createdCategoryProduct2, err := categoryProductService.Create(payload.CreateCategoryProductRequest{
		CategoryName: "Hiburan",
		Description: "Kebutuhan hiburanmu ada disini!.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	err1 := categoryProductService.DeleteById(createdCategoryProduct1.Id)
	if err1 != nil {
		log.Printf("got %v want %v\n", err1.Error(), nil)
	}

	err2 := categoryProductService.DeleteById(createdCategoryProduct2.Id)
	if err2 != nil {
		log.Printf("got %v want %v\n", err2.Error(), nil)
	}
}