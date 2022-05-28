package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	ID           primitive.ObjectID `bson:"_id"`
	Number       *int               `json:"number" validate:"required"`
	AcademicYear *AcademicYear      `json:"academicyear" validate:"required"`
	Year         *int               `json:"year" validate:"required,min=1"`
}
