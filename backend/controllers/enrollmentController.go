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

func AddEnrollment() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "STUDENT"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var enrollment models.Enrollment

		if err := c.BindJSON(&enrollment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user models.User
		var course models.Course

		userid := c.Param("userid")
		courseid := c.Param("courseid")
		real_user_id, _ := primitive.ObjectIDFromHex(userid)
		real_course_id, _ := primitive.ObjectIDFromHex(courseid)
		enrollment.ID = primitive.NewObjectID()

		err := userCollection.FindOne(ctx, bson.M{"_id": real_user_id}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := helpers.MatchUserTypeToUid(c, *user.Username); err != nil {
			c.JSON(http.StatusUnauthorized, "You can only enroll yourself!")
			return
		}

		err = courseCollection.FindOne(ctx, bson.M{"_id": real_course_id}).Decode(&course)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
