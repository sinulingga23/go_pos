package service

import (
	"github.com/sinulingga23/go-pos/payload"
)

type CategoryProductService interface {
	Create(CreateCategoryProductRequest payload.CreateCategoryProductRequest) (*payload.CategoryProduct, error)
}