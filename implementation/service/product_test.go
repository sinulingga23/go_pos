package service

import (
	"context"
	"mime/multipart"
	"net/textproto"
	"os"
	"strings"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/sinulingga23/go-pos/config"
	"github.com/sinulingga23/go-pos/domain"
	"github.com/sinulingga23/go-pos/implementation/repository"
	thirdparty_service "github.com/sinulingga23/go-pos/implementation/service/third-party"
	"github.com/sinulingga23/go-pos/payload"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/option"
)

func TestProductService_Create_Success(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer func() {
		if err := database.Client().Disconnect(ctx); err != nil {
			t.Fatal(err.Error())
		}
	}()

	_, err = database.Collection(repository.ProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = database.Collection(repository.CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		t.Fatal(err.Error())
	}

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("STORAGE_CREDENTIAL_FILE")))
	if err != nil {
		t.Fatal(err.Error())
	}
	defer client.Close()

	wantProductName := "MBP 2022 / RAM 32 GB / SSD 512 GB / M1"
	wantDescription := "For the who want new experiences."
	wantStock := 10
	wantPrice := float64(45000000)

	productRepository := repository.NewProductRepository(database)
	categoryProductRepository := repository.NewCategoryProductRepository(database)
	cloudStorageService := thirdparty_service.NewCloudStorageService(client)
	productService := NewProductService(productRepository, categoryProductRepository, cloudStorageService)

	createdCategoryProduct1, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Elekronik",
		Description: "Menjual alat elektronik.",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	createdCategoryProduct2, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Otomotif",
		Description: "Menjual kebutuhan otomotif.",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	
	images := map[string][]*multipart.FileHeader{}

	images["images"] = []*multipart.FileHeader{
		{Filename: "test",  Header: textproto.MIMEHeader{}, Size: 10},
	}
	
	createdProduct, err := productService.Create(payload.CreateProductRequest{
		CategoryProductIds: []string{
			createdCategoryProduct1.Id.Hex(),
			createdCategoryProduct2.Id.Hex(),
		},
		ProductName: "MBP 2022 / RAM 32 GB / SSD 512 GB / M1",
		Description: "For the who want new experiences.",
		Stock: 10,
		Price: 45000000,
		Images: images,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if strings.Compare(wantProductName, createdProduct.ProductName) != 0 {
		t.Fatalf("got %q want %q\n", createdProduct.ProductName, wantProductName)
	}

	if strings.Compare(wantDescription, createdProduct.Description) != 0 {
		t.Fatalf("got %q want %q\n", createdProduct.Description, wantDescription)
	}

	if wantStock != createdProduct.Stock {
		t.Fatalf("got %v want %v\n", createdProduct.Stock, wantStock)
	}

	if wantPrice != createdProduct.Price {
		t.Fatalf("got %v want %v\n", createdProduct.Price, wantPrice)
	}
}