package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	CategoryProduct struct {
		Id primitive.ObjectID `bson:"_id,omitempty"`
		CategoryName string `bson:"category_name,omitempty"`
		Description string `bson:"description,omitempty"`
	}
)