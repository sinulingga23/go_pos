package repository

import (
	"context"
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
		t.Fatalf(err.Error())
	}
	categoryRepository := NewCategoryProductRepository(database)

	wantCategoryName := "Elektronik"
	wantDescription := "Menjual berbagai macam barang elektronik"
	
	cretaedCategoryProduct, err := categoryRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Elektronik",
		Description: "Menjual berbagai macam barang elektronik",
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	if strings.Compare(wantCategoryName, cretaedCategoryProduct.CategoryName) != 0 {
		t.Fatalf("got %q want %q\n", cretaedCategoryProduct.CategoryName, wantCategoryName)
	}

	if strings.Compare(wantDescription, cretaedCategoryProduct.Description)  != 0 {
		t.Fatalf("got %q want %q\n", cretaedCategoryProduct.Description, wantDescription)
	}
}