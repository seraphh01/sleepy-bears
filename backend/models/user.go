package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Username           *string            `json:"username" validate:"required,min=6,max=32"`
	Email              *string            `json:"email" validate:"email,required"`
	Password           *string            `json:"password" validate:"required,min=8"`
	UserType           *string            `json:"usertype" validate:"required,eq=STUDENT|eq=TEACHER|eq=ADMIN|eq=CHIEF"`
	Name               *string            `json:"name" validate:"required"`
	ProfileDescription *string            `json:"profiledescription"`
	Group              *Group             `json:"group"`
	Token              *string            `json:"token"`
	RefreshToken       *string            `json:"refreshtoken"`
}
