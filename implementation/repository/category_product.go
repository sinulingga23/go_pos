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
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}

	insertedId := result.InsertedID
	singleResult := collection.FindOne(ctx, bson.D{
		bson.E{Key: "_id", Value: insertedId},
	})

	createdCategoryProduct := &domain.CategoryProduct{}
	if err := singleResult.Decode(createdCategoryProduct); err != nil {
		log.Printf("[DATABASE]: %s", err.Error())
		return nil, err
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
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}

	currentCategoryProduct := &domain.CategoryProduct{}
	if err := singleResult.Decode(currentCategoryProduct); err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}

	return currentCategoryProduct, nil
}

func (c *categoryProductRepository) UpdateById(ctx context.Context, id primitive.ObjectID, categoryProduct domain.CategoryProduct) (*domain.CategoryProduct, error) {
	collection := c.database.Collection(CategoryProductCollection)
	_, err := collection.UpdateByID(ctx, id, bson.D{
		bson.E{Key: "$set", Value: categoryProduct},
	})
	if err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}
	
	singleResult := collection.FindOne(ctx, bson.D{
		bson.E{Key: "_id", Value: id},
	})
	if err := singleResult.Err(); err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}

	updatedCategoryProduct := &domain.CategoryProduct{}
	if err := singleResult.Decode(updatedCategoryProduct); err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}

	return updatedCategoryProduct, nil
}

func (c *categoryProductRepository) DeleteById(ctx context.Context, id primitive.ObjectID) (
	int64,
	error,
) {
	collection := c.database.Collection(CategoryProductCollection)
	deleteResult, err := collection.DeleteOne(ctx, bson.D{
		bson.E{Key: "_id", Value: id},
	})
	if err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return 0, err
	}

	return deleteResult.DeletedCount, nil
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
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	categoryProducts := make([]*domain.CategoryProduct, 0)
	for cursor.Next(ctx) {
		currentCategoryProduct := &domain.CategoryProduct{}
		if err := cursor.Decode(currentCategoryProduct); err != nil {
			log.Printf("[DATABASE]: %s\n", err.Error())
			return nil, err
		}
		categoryProducts = append(categoryProducts, currentCategoryProduct)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}

	return categoryProducts, nil
}

func (c *categoryProductRepository) FindAll(ctx context.Context) (
	[]*domain.CategoryProduct,
	error,
) {
	collection := c.database.Collection(CategoryProductCollection)
	cursor, err := collection.Find(ctx, struct {}{})
	if err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	categoryProducts := make([]*domain.CategoryProduct, 0)
	for cursor.Next(ctx) {
		categoryProduct := &domain.CategoryProduct{}
		if err := cursor.Decode(categoryProduct); err != nil {
			log.Printf("[DATABASE]: %s\n", err.Error())
			return nil, err
		}
		categoryProducts = append(categoryProducts, categoryProduct)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}

	return categoryProducts, nil
}
