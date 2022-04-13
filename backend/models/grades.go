package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Grades struct {
	ID     primitive.ObjectID `bson:"_id"`
	User   *User              `json:"user" validate:"required"`
	Course *Course            `json:"course" validate:"required"`
}
