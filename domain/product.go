package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Product struct {
		Id primitive.ObjectID `bson:"_id,omitempty"`
		CategoryProductIds []primitive.ObjectID `bson:"category_product_ids,omitempty"`
		ProductName string `bson:"product_name,omitempty"`
		Description string `bson:"description,omitempty"`
		Stock int `bson:"stock,omitempty"`
		Price float64 `bson:"price,omitempty"`
		UrlImages []UrlImage `bson:"url_images,omitempty"`
	}

	UrlImage struct {
		UrlImageId primitive.ObjectID `bson:"url_image_id,omitempty"`
		Url string `bson:"url,omitempty"`
	}
)