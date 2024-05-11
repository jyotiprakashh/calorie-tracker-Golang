package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Dish      *string             `json:"dish,omitempty"`
	Fat        *float64            `json:"fat,omitempty" `
	Ingredients *string          `json:"ingredients,omitempty" `
	Calories    *string            `json:"calories,omitempty"`
}