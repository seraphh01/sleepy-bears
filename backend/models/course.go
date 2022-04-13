package models

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
  ID         primitive.ObjectID `bson:"_id"`
  Name       *string            `json:"name" validate:"required,min=2,max=40"`
  CourseType *string            `json:"coursetype" validate:"required,eq=MANDATORY|eq=OPTIONAL"`
}
