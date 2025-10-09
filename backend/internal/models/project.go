package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Owner       primitive.ObjectID   `json:"owner" bson:"owner"`
	Members     []primitive.ObjectID `json:"members" bson:"members"`
	CreatedAt   string               `json:"created_at" bson:"created_at"`
	CreatedBy   primitive.ObjectID   `json:"created_by" bson:"created_by"`
}
