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
	ProductCollection = "product"
)

type productRepository struct {
	database *mongo.Database
}

func NewProductRepository(database *mongo.Database) *productRepository {
	return &productRepository{database: database}
}

func (c *productRepository) Create(ctx context.Context, product domain.Product) (
	*domain.Product, 
	error,
) {
	collection := c.database.Collection(ProductCollection)
	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}

	insertedId := result.InsertedID
	singleResult := collection.FindOne(ctx, bson.D{
		bson.E{Key: "_id", Value: insertedId},
	})

	createdProduct := &domain.Product{}
	if err := singleResult.Decode(createdProduct); err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}

	return createdProduct, nil
}

func (c *productRepository) AddUrlImageToProduct(ctx context.Context, id primitive.ObjectID, urlImage domain.UrlImage) error {
	collection := c.database.Collection(ProductCollection)
	_, err := collection.UpdateOne(ctx, bson.D{
		bson.E{Key: "_id", Value: id},
	}, bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{
		Key: "url_images", Value: urlImage,
	}}}})
	if err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return err
	}

	return nil
}

func (c *productRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	collection := c.database.Collection(ProductCollection)
	cursor, err := collection.Find(ctx, struct {}{})
	if err != nil {
		log.Printf("[DATABASE]: %s\n", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	products := make([]*domain.Product, 0)
	for cursor.Next(ctx) {
		product := &domain.Product{}
		if err := cursor.Decode(product); err != nil {
			log.Printf("[DATABASE]: %s\n", err.Error())
			return nil, err
		}

		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}