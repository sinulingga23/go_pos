package service

import (
	"context"
	"log"
	"strings"

	"github.com/sinulingga23/go-pos/api/repository"
	"github.com/sinulingga23/go-pos/definition"
	"github.com/sinulingga23/go-pos/domain"
	"github.com/sinulingga23/go-pos/payload"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type categoryProductService struct {
	categoryProductRepository repository.CategoryProductRepository
}

func NewCategoryProductService(
	categoryProductRepository repository.CategoryProductRepository,
) *categoryProductService {
	return &categoryProductService{categoryProductRepository: categoryProductRepository}
}

func (c *categoryProductService) Create(createCategoryProductRequest payload.CreateCategoryProductRequest) (
	*payload.CategoryProduct,
	error,
) {
	if len(strings.Trim(createCategoryProductRequest.CategoryName, " ")) == 0 {
		return nil, definition.ErrBadRequest
	}

	if len(strings.Trim(createCategoryProductRequest.Description, " ")) == 0 {
		return nil, definition.ErrBadRequest
	}

	categoryProduct, err := c.categoryProductRepository.Create(context.TODO(), domain.CategoryProduct{
		Id: primitive.NewObjectID(),
		CategoryName: createCategoryProductRequest.CategoryName,
		Description: createCategoryProductRequest.Description,
	})
	if err != nil {
		log.Printf("[SERVICE]: %s\n", err.Error())
		return nil, definition.ErrInternalServer
	}

	return &payload.CategoryProduct{
		Id: categoryProduct.Id.String(),
		CategoryName: categoryProduct.CategoryName,
		Description: categoryProduct.Description,
	}, nil
}
