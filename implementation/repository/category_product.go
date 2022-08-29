package repository

import (
	"context"
	"log"

	"github.com/sinulingga23/go-pos/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CategoryProductCollection = `category_product`
)

type categoryProductRepository struct {
	database *mongo.Database
}

func NewCategoryProductRepository(database *mongo.Database) *categoryProductRepository {
	return &categoryProductRepository{database: database}
}

func (c *categoryProductRepository) Create(ctx context.Context, categoryProduct domain.CategoryProduct) (
	*domain.CategoryProduct,
	error,
){
	collection := c.database.Collection(CategoryProductCollection)
	result, err := collection.InsertOne(ctx, categoryProduct)
	if err != nil {
		return nil, err
	}

	insertedId := result.InsertedID
	singleResult := collection.FindOne(ctx, bson.D{
		bson.E{Key: "_id", Value: insertedId},
	})

	createdCategoryProduct := &domain.CategoryProduct{}
	if err := singleResult.Decode(createdCategoryProduct); err != nil {
		log.Printf("[DATABASE]: %s", err.Error())
	}

	return createdCategoryProduct, nil
}

func (c *categoryProductRepository) FindById(ctx context.Context, id primitive.ObjectID) (
	*domain.CategoryProduct,
	error,
) {
	collection := c.database.Collection(CategoryProductCollection)
	singleResult := collection.FindOne(ctx, bson.D{
		bson.E{Key: "_id", Value: id},
	})
	if err := singleResult.Err(); err != nil {
		return nil, err
	}

	currentCategoryProduct := &domain.CategoryProduct{}
	if err := singleResult.Decode(currentCategoryProduct); err != nil {
		return nil, err
	}

	return currentCategoryProduct, nil
}
