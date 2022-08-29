package repository

import (
	"context"
	"errors"
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

func (c *categoryProductRepository) UpdateByID(ctx context.Context, id primitive.ObjectID, categoryProduct domain.CategoryProduct) (
	*domain.CategoryProduct,
	error,	
) {
	collection := c.database.Collection(CategoryProductCollection)
	updateResult, err := collection.UpdateByID(ctx, id, bson.D{
		bson.E{Key: "$set", Value: categoryProduct},
	})
	if err != nil {
		return nil, err
	}
	if updateResult.MatchedCount != 1 {
		return nil, errors.New("There are is a something's wrong.")
	}
	
	singleResult := collection.FindOne(ctx, bson.D{
		bson.E{Key: "_id", Value: id},
	})
	if err := singleResult.Err(); err != nil {
		return nil, err
	}

	updatedCategoryProduct := &domain.CategoryProduct{}
	if err := singleResult.Decode(updatedCategoryProduct); err != nil {
		return nil, err
	}

	return updatedCategoryProduct, nil
}

func (c *categoryProductRepository) DeleteById(ctx context.Context, id primitive.ObjectID) (
	error,
) {
	collection := c.database.Collection(CategoryProductCollection)
	deleteResult, err := collection.DeleteOne(ctx, bson.D{
		bson.E{Key: "_id", Value: id},
	})
	if err != nil {
		return err
	}
	
	if deleteResult.DeletedCount != 1 {
		return errors.New("There are is a something's wrong.")
	}

	return nil
}

func (c *categoryProductRepository) FindByIds(ctx context.Context, ids []primitive.ObjectID) (
	[]*domain.CategoryProduct,
	error,
) {
	collection := c.database.Collection(CategoryProductCollection)
	cursor, err := collection.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	categoryProducts := make([]*domain.CategoryProduct, 0)
	for cursor.Next(ctx) {
		currentCategoryProduct := &domain.CategoryProduct{}
		if err := cursor.Decode(currentCategoryProduct); err != nil {
			return nil, err
		}
		categoryProducts = append(categoryProducts, currentCategoryProduct)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categoryProducts, nil
}
