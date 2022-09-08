package repository

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/sinulingga23/go-pos/config"
	"github.com/sinulingga23/go-pos/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
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