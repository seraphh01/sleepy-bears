package models

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
  ID                 primitive.ObjectID `bson:"_id"`
  Username           *string            `json:"Username" validate:"required,min=8,max=32"`
  Email              *string            `json:"Email" validate:"email,required"`
  Password           *string            `json:"Password" validate:"required,min=8,max=20"`
  Name               *string            `json:"Name" validate:"required"`
  PhoneNumber        *string            `json:"PhoneNumber" validate:"required"`
  Role               *string            `json:"Role"`
  ProfileDescription *string            `json:"ProfileDescription"`
}
