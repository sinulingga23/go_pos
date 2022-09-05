package repository

import (
	"context"

	"github.com/sinulingga23/go-pos/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryProductRepository interface {
	Create(ctx context.Context, categoryProduct domain.CategoryProduct) (*domain.CategoryProduct, error)
	FindById(ctx context.Context, id primitive.ObjectID) (*domain.CategoryProduct, error)
	UpdateById(ctx context.Context, id primitive.ObjectID, categoryProduct domain.CategoryProduct) (*domain.CategoryProduct, error)
	DeleteById(ctx context.Context, id primitive.ObjectID) (int64, error)
	FindByIds(ctx context.Context, ids []primitive.ObjectID) ([]*domain.CategoryProduct, error)
	FindAll(ctx context.Context) ([]*domain.CategoryProduct, error)
}