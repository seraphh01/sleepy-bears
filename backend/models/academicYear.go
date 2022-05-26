package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AcademicYear struct {
	ID        primitive.ObjectID `bson:"_id"`
	StartDate primitive.DateTime `json:"start_date" validate:"required"`
	EndDate   primitive.DateTime `json:"end_date" validate:"required"`
}
