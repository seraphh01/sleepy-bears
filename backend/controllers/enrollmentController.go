package controllers

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
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
		err = courseCollection.FindOne(ctx, bson.M{"_id": real_course_id}).Decode(&course)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		enrollment.User = &user
		enrollment.Course = &course
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
