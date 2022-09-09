package repository

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/sinulingga23/go-pos/config"
	"github.com/sinulingga23/go-pos/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func TestCategoryProductRepository_Create_Success(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
	
	_, err = database.Collection(CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	categoryProductRepository := NewCategoryProductRepository(database)

	wantCategoryName := "Elektronik"
	wantDescription := "Menjual berbagai macam barang elektronik"
	
	cretaedCategoryProduct, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Elektronik",
		Description: "Menjual berbagai macam barang elektronik",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if strings.Compare(wantCategoryName, cretaedCategoryProduct.CategoryName) != 0 {
		t.Fatalf("got %q want %q\n", cretaedCategoryProduct.CategoryName, wantCategoryName)
	}

	if strings.Compare(wantDescription, cretaedCategoryProduct.Description)  != 0 {
		t.Fatalf("got %q want %q\n", cretaedCategoryProduct.Description, wantDescription)
	}
}

func TestCategoryProductRepository_FindById_Exists(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	_, err = database.Collection(CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	categoryProductRepository := NewCategoryProductRepository(database)

	wantId := primitive.NewObjectID()
	wantCategoryName := "Kesehatan"
	wantDescription := "Menyediakan berbagai produk kesehatan."

	createdCategoryProduct, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: wantId,
		CategoryName: "Kesehatan",
		Description: "Menyediakan berbagai produk kesehatan.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	currentCategoryProduct, err := categoryProductRepository.FindById(ctx, createdCategoryProduct.Id)
	if err != nil {
		log.Fatal(err.Error())
	}

	if wantId != currentCategoryProduct.Id {
		log.Fatalf("got %q want %q\n", currentCategoryProduct.Id, wantId)
	}

	if strings.Compare(wantCategoryName, currentCategoryProduct.CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", currentCategoryProduct.CategoryName, wantCategoryName)
	}

	if strings.Compare(wantDescription, currentCategoryProduct.Description) != 0 {
		log.Fatalf("got %q want %q\n", wantDescription, currentCategoryProduct.Description)
	}
}

func TestCategoryProductRepository_FindById_NotExists(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantError1 := mongo.ErrNoDocuments
	wantError2 := mongo.ErrNoDocuments

	categoryProductRepository := NewCategoryProductRepository(database)

	_, err = categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Hiburan",
		Description: "Semua kebutuhan hiburanmu ada disini!.",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, err1 := categoryProductRepository.FindById(ctx, primitive.NewObjectID())
	_, err2 := categoryProductRepository.FindById(ctx, primitive.ObjectID{})

	if wantError1 != err1 {
		log.Fatalf("got %q want %q\n", err1.Error(), wantError1.Error())
	}

	if wantError2 != err2 {
		log.Fatalf("got %q want %q\n", err2.Error(), wantError2.Error())
	}
}

func TestCategoryProductRepository_UpdateById_Success(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	_, err = database.Collection(CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantId := primitive.NewObjectID()
	wantCategoryName := "Fashion"
	wantDescription := "Semua kebutuhan fashionmu ada disini!."

	categoryProductRepository := NewCategoryProductRepository(database)


	createdCategoryProduct, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: wantId,
		CategoryName: "Pakaian",
		Description: "Semua kebutuhan pakaian ada di sini!.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	updatedCategoryProduct, err := categoryProductRepository.UpdateById(ctx, createdCategoryProduct.Id, domain.CategoryProduct{
		Id: createdCategoryProduct.Id,
		CategoryName: "Fashion",
		Description: "Semua kebutuhan fashionmu ada disini!.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	if strings.Compare(wantId.String(), updatedCategoryProduct.Id.String()) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct.Id.String(), wantId.String())
	}

	if strings.Compare(wantCategoryName, updatedCategoryProduct.CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct.CategoryName, wantCategoryName)
	}

	if strings.Compare(wantDescription, updatedCategoryProduct.Description) != 0 {
		log.Fatalf("got %q want %q\n", updatedCategoryProduct.Description, wantDescription)
	}
}

func TestCategoryProductRepository_UpdateById_NotExists(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantError1 := mongo.ErrNoDocuments
	wantError2 := mongo.ErrNoDocuments

	categoryProductRepository := NewCategoryProductRepository(database)

	createdCategoryProduct, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Pakaian",
		Description: "Semua kebutuhan pakaian ada di sini!.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err1 := categoryProductRepository.UpdateById(ctx, primitive.NewObjectID(), domain.CategoryProduct{
		Id: createdCategoryProduct.Id,
		CategoryName: "Fashion",
		Description: "Semua kebutuhan fashionmu ada disini!.",
	})

	_, err2 := categoryProductRepository.UpdateById(ctx, primitive.ObjectID{}, domain.CategoryProduct{
		Id: createdCategoryProduct.Id,
		CategoryName: "Fashion",
		Description: "Semua kebutuhan fashionmu ada disini!.",
	})

	if wantError1 != err1 {
		log.Fatalf("got %q want %q\n", err1.Error(), wantError1.Error())
	}

	if wantError2 != err2 {
		log.Fatalf("got %q want %q\n", err2.Error(), wantError2.Error())
	}
}
