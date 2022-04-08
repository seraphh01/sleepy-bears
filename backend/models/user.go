package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Username           *string            `json:"Username" validate:"required,min=8,max=32"`
	Email              *string            `json:"Email" validate:"email,required"`
	Password           *string            `json:"Password" validate:"required,min=8,max=20"`
	UserType           *string            `json:"UserType" validate:"required,eq=STUDENT|eq=TEACHER|eq=ADMIN|eq=CHIEF"`
	Name               *string            `json:"Name" validate:"required"`
	ProfileDescription *string            `json:"ProfileDescription"`
	Token              *string            `json:"Token"`
	Refresh_token      *string            `json:"Refresh_token"`
	User_id            string             `json:"User_id"`
}
