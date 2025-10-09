package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	State       string             `json:"state" bson:"state"`
	Project     primitive.ObjectID `json:"project" bson:"project"`
	Assignee    primitive.ObjectID `json:"assignee" bson:"assignee"`
	Number      int                `json:"number" bson:"number"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	CreatedBy   primitive.ObjectID `json:"created_by" bson:"created_by"`
}

type NextTaskNumber struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Number  int                `json:"number" bson:"number"`
	Project primitive.ObjectID `json:"project" bson:"project"`
}
