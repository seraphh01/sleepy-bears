package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Grade struct {
	Grade int `json:"grade" validate:"required,min=0,max=10"`
}

type Enrollment struct {
	ID     primitive.ObjectID `bson:"_id"`
	User   *User              `json:"user" validate:"required"`
	Course *Course            `json:"course" validate:"required"`
	Grades []Grade            `json:"grades"`
}
