package service

import "github.com/sinulingga23/go-pos/payload"

type ProductService interface {
	Create(createProduct payload.CreateProductRequest) (*payload.Product, error)
	AddImagesToProduct(productId string, addImageToProduct payload.AddImageToProductRequest) error
}