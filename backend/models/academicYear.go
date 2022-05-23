package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AcademicYear struct {
	ID        primitive.ObjectID `bson:"_id"`
	StartDate time.Time          `json:"start_date" validate:"required"`
	EndDate   time.Time          `json:"end_date" validate:"required"`
}
