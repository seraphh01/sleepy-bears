package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AcademicYear struct {
	ID        primitive.ObjectID `bson:"_id"`
	StartDate primitive.DateTime `json:"startdate" validate:"required"`
	EndDate   primitive.DateTime `json:"enddate" validate:"required"`
}
