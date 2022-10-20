package repository

import (
	"context"

	"github.com/sinulingga23/go-pos/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) (*domain.Product, error)
	AddUrlImageToProduct(ctx context.Context, id primitive.ObjectID, imageUrl domain.UrlImage) error
	FindAll(ctx context.Context) ([]*domain.Product, error)
}