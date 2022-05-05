package models

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Amount struct {
  Max int `json:"max" validate:"required,min=20"`
}

type Course struct {
  ID         primitive.ObjectID `bson:"_id"`
  Name       *string            `json:"name" validate:"required,min=2,max=40"`
  CourseType *string            `json:"coursetype" validate:"required,eq=MANDATORY|eq=OPTIONAL"`
  Year       *int               `json:"year" validate:"required,min=1"`
  Proposer   *User              `json:"proposer" validate:"required"`
  MaxAmount  *Amount            `json:"maxamount" validate:"required,min=20"`
}
