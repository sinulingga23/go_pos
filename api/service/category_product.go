package service

import (
	"github.com/sinulingga23/go-pos/payload"
)

type CategoryProductService interface {
	Create(createCategoryProductRequest payload.CreateCategoryProductRequest) (*payload.CategoryProduct, error)
	FindById(id string) (*payload.CategoryProduct, error)
	UpdateById(id string, updateCategoryProductRequest payload.UpdateCategoryProductRequest) (*payload.CategoryProduct, error)
}