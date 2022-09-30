package handler

import (
	"github.com/sinulingga23/go-pos/api/service"
)


type handler struct {
	categoryProductService service.CategoryProductService
	productService service.ProductService
}

func NewHandler(
	categoryProductService service.CategoryProductService,
	productService service.ProductService) *handler {
	return &handler{
		categoryProductService: categoryProductService,
		productService: productService,
	}
}