package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Amount struct {
	Max int `json:"max" validate:"required,min=20"`
}

type Course struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         *string            `json:"name" validate:"required,min=2,max=40"`
	CourseType   *string            `json:"coursetype" validate:"required,eq=MANDATORY|eq=OPTIONAL"`
	AcademicYear *AcademicYear      `json:"academicyear" validate:"required"`
	YearOfStudy  *int               `json:"yearofstudy" validate:"required,min=1"`
	Proposer     *User              `json:"proposer" validate:"required"`
	MaxAmount    *Amount            `json:"maxamount" validate:"required,min=20"`
	Credits      *int               `json:"credits" validate:"required,min=2"`
	Semester     *int               `json:"semester" validate:"required,min=1,max=2"`
}
