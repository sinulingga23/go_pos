package payload

type (
	CategoryProduct struct {
		Id string `json:"id"`
		CategoryName string `json:"categoryName"`
		Description string `json:"description"`
	}

	CreateCategoryProductRequest struct {
		CategoryName string `json:"categoryName"`
		Description string `json:"description"`
	}

	UpdateCategoryProductRequest struct {
		Id string `json:"id"`
		CategoryName string `json:"categoryName"`
		Description string `json:"description"`
	}
)