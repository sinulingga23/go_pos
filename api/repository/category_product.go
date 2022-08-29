package repository

import (
	"context"

	"github.com/sinulingga23/go-pos/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryProductRepository interface {
	Create(ctx context.Context, categoryProduct domain.CategoryProduct) (*domain.CategoryProduct, error)
	FindById(ctx context.Context, id primitive.ObjectID) (*domain.CategoryProduct, error)
}