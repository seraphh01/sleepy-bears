package models

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO list of ints
type Enrollment struct {
  ID     primitive.ObjectID `bson:"_id"`
  User   *User              `json:"user" validate:"required"`
  Course *Course            `json:"course" validate:"required"`
  Grade  *int               `json:"value" validate:"required,min=0,max=10"`
}
