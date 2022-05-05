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

//TODO add enrollment function into courseCollection. Make sure user doesn't enroll in the same course twice.
//TODO 2 when student enrolls, check that the maximum number of enrollments at that course was not reached
func AddEnrollment() gin.HandlerFunc {
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

		userid := user.ID.Hex()
		real_user_id, _ := primitive.ObjectIDFromHex(userid)
		courseid := c.Param("courseid")
		real_course_id, _ := primitive.ObjectIDFromHex(courseid)
		enrollment.ID = primitive.NewObjectID()

		err = proposedCourseCollection.FindOne(ctx, bson.M{"_id": real_course_id}).Decode(&course)
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

func GradeStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "TEACHER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var enrollment models.Enrollment
		var grade models.Grade
		username := c.Param("studentusername")

		courseid := c.Param("courseid")
		realCourseID, _ := primitive.ObjectIDFromHex(courseid)

		err := enrollmentCollection.FindOne(ctx, bson.M{"user.username": username, "course._id": realCourseID}).Decode(&enrollment)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.BindJSON(&grade); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if helpers.MatchUserToUsername(c, *enrollment.Course.Proposer.Username) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You can only grade your own course!"})
			return
		}
		enrollment.Grades = append(enrollment.Grades, grade)

		update := bson.M{"grades": enrollment.Grades}
		result, err := enrollmentCollection.UpdateOne(ctx, bson.M{"user.username": username, "course._id": realCourseID}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var updatedGrades models.Enrollment
		if result.MatchedCount == 1 {
			err := enrollmentCollection.FindOne(ctx, bson.M{"user.username": username, "course._id": realCourseID}).Decode(&updatedGrades)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, updatedGrades)
	}
}

func ViewGrades() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "STUDENT"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var username = c.GetString("username")

		courseid := c.Param("courseid")
		real_course_id, _ := primitive.ObjectIDFromHex(courseid)
		var enrollment models.Enrollment
		err := enrollmentCollection.FindOne(ctx, bson.M{"user.username": username, "course._id": real_course_id}).Decode(&enrollment)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, enrollment.Grades)
	}
}

func DeleteEnrollmentsByCourseID(courseId primitive.ObjectID) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	_, err := enrollmentCollection.DeleteMany(ctx, bson.M{"course._id": courseId})
	if err != nil {
		return err
	}
	return nil
}
