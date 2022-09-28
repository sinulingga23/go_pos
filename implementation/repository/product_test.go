package repository

import (
	"context"
	"strings"
	"testing"

	"github.com/sinulingga23/go-pos/config"
	"github.com/sinulingga23/go-pos/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func TestProductRepository_Create_Success(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.Collection(ProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		t.Fatal(err.Error())
	}

	categoryProductRepository := NewCategoryProductRepository(database)
	productRepository := NewProductRepository(database)

	createdCategoryProduct1, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Hiburan",
		Description: "Menyediakan berbagai kebutuhan hiburan.",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	createdCategoryProduct2, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Elektronik",
		Description: "Menyediakan berbagai macam alat elektronik.",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	
	wantCategoryProductIds := []primitive.ObjectID{
		createdCategoryProduct1.Id,
		createdCategoryProduct2.Id,
	}
	wantProductName := "Macbook Pro 2022 / RAM 32 GB/ SSD 512 GB"
	wantDescription := "Macbook Pro yang dapat digunakan untuk keperluan apapun."
	wantStock := 23
	wantPrice := float64(45000000)
	wantUrlImages := []domain.UrlImage{}

	createdProduct, err := productRepository.Create(ctx, domain.Product{
		Id: primitive.NewObjectID(),
		CategoryProductIds: []primitive.ObjectID{
			createdCategoryProduct1.Id,
			createdCategoryProduct2.Id,
		},
		ProductName: "Macbook Pro 2022 / RAM 32 GB/ SSD 512 GB",
		Description: "Macbook Pro yang dapat digunakan untuk keperluan apapun.",
		Stock: 23,
		Price: 45000000,
		UrlImages: []domain.UrlImage{},
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(wantCategoryProductIds) != len(createdProduct.CategoryProductIds) {
		t.Fatalf("got %q want %q\n", len(createdProduct.CategoryProductIds), len(wantCategoryProductIds))
	}

	for i := 0; i < len(wantCategoryProductIds); i++ {
		if strings.Compare(wantCategoryProductIds[i].String(), createdProduct.CategoryProductIds[i].String()) != 0 {
			t.Fatalf("got %q want %q\n", createdProduct.CategoryProductIds[i].String(), wantCategoryProductIds[i].String())
		}
	}

	if strings.Compare(wantProductName, createdProduct.ProductName) != 0 {
		t.Fatalf("got %q want %q\n", createdProduct.ProductName, wantProductName)
	}

	if strings.Compare(wantDescription, createdProduct.Description) != 0 {
		t.Fatalf("got %q want %q\n", createdProduct.Description, wantDescription)
	}

	if wantStock != createdProduct.Stock  {
		t.Fatalf("got %q want %q\n", createdProduct.Stock, wantStock)
	}

	if wantPrice != createdProduct.Price  {
		t.Fatalf("got %v want %v\n", createdProduct.Price, wantPrice)
	}

	if len(wantUrlImages) != len(createdProduct.UrlImages) {
		t.Fatalf("got %v want %v\n", len(createdProduct.UrlImages), len(wantUrlImages))
	}
}