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
	"go.mongodb.org/mongo-driver/mongo"
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

func (c *categoryProductService) FindById(id string) (
	*payload.CategoryProduct, 
	error,
) {
	idOID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, definition.ErrBadRequest
	}

	categoryProduct, err := c.categoryProductRepository.FindById(context.TODO(), idOID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, definition.ErrDataNotFound
		}

		return nil, definition.ErrInternalServer
	}

	return &payload.CategoryProduct{
		Id: categoryProduct.Id.String(),
		CategoryName: categoryProduct.CategoryName,
		Description: categoryProduct.Description,
	}, nil
}


func (c *categoryProductService) UpdateById(id string, updateCategoryProductRequest payload.UpdateCategoryProductRequest) (
	*payload.CategoryProduct,
	error,
) {
	if strings.Compare(id, updateCategoryProductRequest.Id) != 0 {
		return nil, definition.ErrBadRequest
	}

	if len(strings.Trim(updateCategoryProductRequest.CategoryName, " ")) == 0 {
		return nil, definition.ErrBadRequest
	}

	if len(strings.Trim(updateCategoryProductRequest.Description, " ")) == 0 {
		return nil, definition.ErrBadRequest
	}

	idOID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, definition.ErrBadRequest
	}

	// TODO: Transaction
	currentCategoryProduct, err := c.categoryProductRepository.FindById(context.TODO(), idOID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, definition.ErrDataNotFound
		}

		return nil, definition.ErrInternalServer
	}

	currentCategoryProduct.CategoryName = updateCategoryProductRequest.CategoryName
	currentCategoryProduct.Description = updateCategoryProductRequest.Description
	
	updatedCategoryProduct, err := c.categoryProductRepository.UpdateById(context.TODO(), idOID, *currentCategoryProduct)
	if err != nil {
		return nil, definition.ErrInternalServer
	}

	return &payload.CategoryProduct{
		Id: updatedCategoryProduct.Id.String(),
		CategoryName: updateCategoryProductRequest.CategoryName,
		Description: updateCategoryProductRequest.Description,
	}, nil
}

func (c *categoryProductService) DeleteById(id string) error {
	idOID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return definition.ErrBadRequest
	}
	
	deletedCount, err := c.categoryProductRepository.DeleteById(context.TODO(), idOID)
	if err != nil {
		return definition.ErrInternalServer
	}

	if deletedCount == 0 {
		return definition.ErrDataNotFound
	}

	return nil
}

func (c *categoryProductService) FindAll() ([]*payload.CategoryProduct, error) {
	categoryProducts, err := c.categoryProductRepository.FindAll(context.TODO())
	if err != nil {
		return nil, err
	}

	result := make([]*payload.CategoryProduct, 0)
	for _,categoryProduct := range categoryProducts {
		result = append(result, &payload.CategoryProduct{
			Id: categoryProduct.Id.String(),
			CategoryName: categoryProduct.CategoryName,
			Description: categoryProduct.Description,
		})
	}

	if len(result) == 0 {
		return nil, definition.ErrDataNotFound
	}

	return result, nil
}
