package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Email              *string            `json:"email"`
	Name               *string            `json:"name"`
	PhoneNumber        *string            `json:"PhoneNumber"`
	Enrollments        *[]string          `json:"Enrollments"` // TODO: change from array of strings to array of enrollments
	Role               *string            `json:"Role"`
	ProfileDescription *string            `json:"ProfileDescription"`
}
