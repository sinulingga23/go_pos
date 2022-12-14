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

func TestCategoryProductRepository_DeleteById(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	var wantDeletedCount1 int64 = 1
	var wantDeletedCount2 int64 = 1

	categoryProductRepository := NewCategoryProductRepository(database)

	createdCategoryProduct1, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Pinjaman",
		Description: "Bayar semua pinjamanmu di sini!.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	createdCategoryProduct2, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Otomotif",
		Description: "Tempat yang tepat untuk mencari kebutuhan otomotifmu.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	deletedCount1, err := categoryProductRepository.DeleteById(ctx, createdCategoryProduct1.Id)
	if err != nil {
		log.Fatal(err.Error())
	}

	deletedCount2, err := categoryProductRepository.DeleteById(ctx, createdCategoryProduct2.Id)
	if err != nil {
		log.Fatal(err.Error())
	}

	if wantDeletedCount1 != deletedCount1 {
		log.Fatalf("got %q want %q\n", deletedCount1, wantDeletedCount1)
	}

	if wantDeletedCount2 != deletedCount2 {
		log.Fatalf("got %q want %q\n", deletedCount2, wantDeletedCount2)
	}
}


func TestCategoryProductRepository_FindByIds_Success(t *testing.T) {
	ctx := context.TODO()
	database, err := config.ConnectToMongoDb(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = database.Collection(CategoryProductCollection).DeleteMany(ctx, struct {}{})
	if err != nil {
		log.Fatal(err.Error())
	}

	wantLengthData := 5
	wantCategoryName1 := "Pakaian"
	wantDescription1 := "Semua kebutuhan pakaian ada di sini!."

	wantCategoryName3 := "Otomotif"
	wantDescription3 := "Tempat yang tepat untuk mencari kebutuhan otomotifmu."

	wantCategoryName5 := "Kesehatan"
	wantDescription5 := "Menyediakan berbagai produk kesehatan."

	categoryProductRepository := NewCategoryProductRepository(database)

	createdCategoryProduct1, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Pakaian",
		Description: "Semua kebutuhan pakaian ada di sini!.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	createdCategoryProduct2, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		CategoryName: "Fashion",
		Description: "Semua kebutuhan fashionmu ada disini!.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	createdCategoryProduct3, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Otomotif",
		Description: "Tempat yang tepat untuk mencari kebutuhan otomotifmu.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	createdCategoryProduct4, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Pinjaman",
		Description: "Bayar semua pinjamanmu di sini!.",
	})

	createdCategoryProduct5, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Kesehatan",
		Description: "Menyediakan berbagai produk kesehatan.",
	})

	categoryProducts, err := categoryProductRepository.FindByIds(ctx, []primitive.ObjectID{
		createdCategoryProduct1.Id,
		createdCategoryProduct2.Id,
		createdCategoryProduct3.Id,
		createdCategoryProduct4.Id,
		createdCategoryProduct5.Id,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	if wantLengthData != len(categoryProducts) {
		log.Fatalf("got %q want %q\n", len(categoryProducts), wantLengthData)
	}

	if strings.Compare(wantCategoryName1, categoryProducts[0].CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", categoryProducts[0].CategoryName, wantCategoryName1)
	}

	if strings.Compare(wantDescription1, categoryProducts[0].Description) != 0 {
		log.Fatalf("got %q want %q\n", categoryProducts[0].Description, wantDescription1)
	}

	if strings.Compare(wantCategoryName3, categoryProducts[2].CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", categoryProducts[2].CategoryName, wantCategoryName3)
	}

	if strings.Compare(wantDescription3, categoryProducts[2].Description) != 0 {
		log.Fatalf("got %q want %q\n", categoryProducts[2].Description, wantDescription3)
	}

	if strings.Compare(wantCategoryName5, categoryProducts[4].CategoryName) != 0 {
		log.Fatalf("got %q want %q\n", categoryProducts[4].CategoryName, wantCategoryName5)
	}

	if strings.Compare(wantDescription5, categoryProducts[4].Description) != 0 {
		log.Fatalf("got %q want %q\n", categoryProducts[4].Description, wantDescription5)
	}
}

func TestCategoryProductRepository_FindAll_Success(t *testing.T) {
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

	createdCategoryProduct1, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Pakaian",
		Description: "Semua kebutuhan pakaian ada di sini!.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	createdCategoryProduct2, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		CategoryName: "Fashion",
		Description: "Semua kebutuhan fashionmu ada disini!.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	createdCategoryProduct3, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Otomotif",
		Description: "Tempat yang tepat untuk mencari kebutuhan otomotifmu.",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	createdCategoryProduct4, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Pinjaman",
		Description: "Bayar semua pinjamanmu di sini!.",
	})

	createdCategoryProduct5, err := categoryProductRepository.Create(ctx, domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: "Kesehatan",
		Description: "Menyediakan berbagai produk kesehatan.",
	})

	wantCategoryProducts := []*domain.CategoryProduct{
		createdCategoryProduct1,
		createdCategoryProduct2,
		createdCategoryProduct3, 
		createdCategoryProduct4,
		createdCategoryProduct5,
	}
	wantLengthData := len(wantCategoryProducts)

	cetegoryProducts, err := categoryProductRepository.FindAll(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	if wantLengthData != len(cetegoryProducts) {
		log.Fatalf("got %q want %q\n", len(cetegoryProducts), wantLengthData)
	}

	for i := 0; i < len(cetegoryProducts); i++ {
		if strings.Compare(wantCategoryProducts[i].Id.String(), cetegoryProducts[i].Id.String()) != 0 {
			log.Fatalf("got %q want %q\n", cetegoryProducts[i].Id.String(), wantCategoryProducts[i].Id.String())
		}
	}
}