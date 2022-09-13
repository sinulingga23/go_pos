package handler

import (
	"github.com/sinulingga23/go-pos/api/service"
)


type handler struct {
	categoryProductService service.CategoryProductService
}

func NewHandler(categoryProductService service.CategoryProductService) *handler {
	return &handler{
		categoryProductService: categoryProductService,
	}
}