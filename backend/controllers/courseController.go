package controllers

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var courseCollection *mongo.Collection = database.OpenCollection(database.Client, "Course")
var proposedCourseCollection *mongo.Collection = database.OpenCollection(database.Client, "ProposedCourse")

func ApproveCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "CHIEF"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var course models.Course

		courseID := c.Param("id")
		realCourseID, _ := primitive.ObjectIDFromHex(courseID)
		var studentEnrollmentCount = GetEnrollmentsCountByCourseID(c, realCourseID)

		if studentEnrollmentCount < 0 {
			return
		}

		if studentEnrollmentCount < 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot approve the course since less than 20 students are enrolled"})
			return
		}
		count, err := proposedCourseCollection.CountDocuments(ctx, bson.M{"_id": realCourseID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This course is not in the proposed courses list!"})
			return
		}

		err = proposedCourseCollection.FindOne(ctx, bson.M{"_id": realCourseID}).Decode(&course)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(course)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		result, err := proposedCourseCollection.DeleteOne(ctx, bson.M{"_id": realCourseID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		course.ID = primitive.NewObjectID()

		resultInsertionNumber, err := courseCollection.InsertOne(ctx, course)
		if err != nil {
			msg := fmt.Sprintf("Course item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func GetCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var courses []models.Course
		cursor, err := courseCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for cursor.Next(ctx) {
			var course models.Course
			err := cursor.Decode(&course)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			courses = append(courses, course)
		}
		if len(courses) > 0 {
			c.JSON(http.StatusOK, courses)
		} else {
			c.JSON(http.StatusOK, "No courses available!")
		}
	}
}

func AddProposedCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "TEACHER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var course models.Course
		var user models.User

		username := c.GetString("username")
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := helpers.MatchUserTypeToUid(c, *user.Username); err != nil {
			c.JSON(http.StatusUnauthorized, "You can only propose your own courses!")
			return
		}
		realUserID, _ := primitive.ObjectIDFromHex(user.ID.Hex())

		count, err := proposedCourseCollection.CountDocuments(ctx, bson.M{"proposer._id": realUserID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count >= 2 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This teacher has already proposed two courses!"})
			return
		}

		countAccepted, err := courseCollection.CountDocuments(ctx, bson.M{"proposer._id": realUserID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if countAccepted+count >= 2 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This teacher has already proposed two courses!"})
			return
		}

		ctype := "OPTIONAL"
		course.CourseType = &ctype
		course.Proposer = &user
		if err := c.BindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(course)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		course.ID = primitive.NewObjectID()

		resultInsertionNumber, insertErr := proposedCourseCollection.InsertOne(ctx, course)
		if insertErr != nil {
			msg := fmt.Sprintf("Course item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}

func GetProposedCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var courses []models.Course
		cursor, err := proposedCourseCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for cursor.Next(ctx) {
			var course models.Course
			err := cursor.Decode(&course)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			courses = append(courses, course)
		}
		if len(courses) > 0 {
			c.JSON(http.StatusOK, courses)
		} else {
			c.JSON(http.StatusOK, "No proposed courses available")
		}
	}
}

func GetCoursesByYear() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		year, err := strconv.Atoi(c.Param("year"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var courses []models.Course
		cursor, err := courseCollection.Find(ctx, bson.M{"year": year})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for cursor.Next(ctx) {
			var course models.Course
			err := cursor.Decode(&course)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			courses = append(courses, course)
		}
		if len(courses) > 0 {
			c.JSON(http.StatusOK, courses)
		} else {
			c.JSON(http.StatusOK, "No courses available in this academic year")
		}
	}
}
