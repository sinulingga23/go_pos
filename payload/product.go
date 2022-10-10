package payload

import "mime/multipart"


type (
	Product struct {
		Id string `json:"id"`
		CategoryProducts []CategoryProduct `json:"categoryProducts"`
		ProductName string `json:"productName"`
		Description string `json:"description"`
		Stock int `json:"stock"`
		Price float64 `json:"price"`
		UrlImages []UrlImage `json:"urlImages"`
	}
	
	UrlImage struct {
		UrlImageId string `json:"urlImageId"`
		Url string `json:"url"`
	}

	CreateProductRequest struct {
		CategoryProductIds []string
		ProductName string
		Description string
		Stock int
		Price float64
		Images map[string][]*multipart.FileHeader
	}

	AddImageToProductRequest struct {
		ProudctId string
		Images map[string][]*multipart.FileHeader
	}
)