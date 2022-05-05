package controllers

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var enrollmentCollection *mongo.Collection = database.OpenCollection(database.Client, "Enrollment")

// Used by a student to enroll to an optional course
func AddOptionalEnrollment() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "STUDENT"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var course models.Course
		var enrollment models.Enrollment
		username := c.GetString("username")
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := helpers.MatchUserTypeToUid(c, *user.Username); err != nil {
			c.JSON(http.StatusUnauthorized, "You can only enroll yourself!")
			return
		}
		courseid := c.Param("courseid")
		userid := user.ID.Hex()
		real_user_id, _ := primitive.ObjectIDFromHex(userid)
		real_course_id, _ := primitive.ObjectIDFromHex(courseid)
		enrollment.ID = primitive.NewObjectID()

		err = proposedCourseCollection.FindOne(ctx, bson.M{"_id": real_course_id}).Decode(&course)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		allowedType := "OPTIONAL"
		if *course.CourseType != allowedType {
			c.JSON(http.StatusBadRequest, "You can only enroll to optional courses!")
			return
		}

		enrollment.User = &user
		enrollment.Course = &course
		count, err := enrollmentCollection.CountDocuments(ctx, bson.M{"user._id": real_user_id, "course._id": real_course_id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count != 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "You have already enrolled for this course!"})
			return
		}
		resultInsertionNumber, insertErr := enrollmentCollection.InsertOne(ctx, enrollment)
		if insertErr != nil {
			msg := fmt.Sprintf("Enrollment item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}

func GetEnrollmentsCountByCourseID(c *gin.Context, courseID primitive.ObjectID) int64 {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	count, err := enrollmentCollection.CountDocuments(ctx, bson.M{"course._id": courseID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return -1
	}
	return count
}
