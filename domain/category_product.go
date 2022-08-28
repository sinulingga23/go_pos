package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

var (
	CategoryProduct struct {
		Id primitive.ObjectID `bson:"_id"`
		CategoryName string `bson:"category_name,omitempty"`
		Description string `bson:"description,omitempty"`
	}
)