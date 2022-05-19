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
)

var enrollmentCollection = database.OpenCollection(database.Client, "Enrollment")

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
		realUserId, _ := primitive.ObjectIDFromHex(userid)
		courseid := c.Param("courseid")
		realCourseId, _ := primitive.ObjectIDFromHex(courseid)
		enrollment.ID = primitive.NewObjectID()

		err = proposedCourseCollection.FindOne(ctx, bson.M{"_id": realCourseId}).Decode(&course)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		enrollment.User = &user
		enrollment.Course = &course
		count, err := enrollmentCollection.CountDocuments(ctx, bson.M{"user._id": realUserId, "course._id": realCourseId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count != 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "You have already enrolled for this course!"})
			return
		}
		var studentEnrollmentCount = int(GetEnrollmentsCountByCourseID(c, realCourseId))
		if studentEnrollmentCount >= course.MaxAmount.Max {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "You have already reached the maximum enrollment for this course!"})
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
		realCourseId, _ := primitive.ObjectIDFromHex(courseid)
		var enrollment models.Enrollment
		err := enrollmentCollection.FindOne(ctx, bson.M{"user.username": username, "course._id": realCourseId}).Decode(&enrollment)
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

func GetOptionalEnrollmentsByStudentUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		username := c.Param("username")
		var courses []models.Course
		cursor, err := enrollmentCollection.Find(ctx, bson.M{"user.username": username, "course.coursetype": "OPTIONAL"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for cursor.Next(ctx) {
			var enrollment models.Enrollment
			err := cursor.Decode(&enrollment)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			courses = append(courses, *enrollment.Course)
		}
		if len(courses) > 0 {
			c.JSON(http.StatusOK, courses)
		} else {
			c.JSON(http.StatusOK, "No courses available for this student")
		}
	}
}

func GetAverageGradeByCourseID(c *gin.Context, realCourseId primitive.ObjectID) float64 {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	cursor, err := enrollmentCollection.Find(ctx, bson.M{"course._id": realCourseId})

	if err != nil {
		fmt.Println(err)
		return -1
	}
	var total float64 = 0
	var totalCount = 0
	for cursor.Next(ctx) {
		var enrollment models.Enrollment
		err := cursor.Decode(&enrollment)
		if err != nil {
			fmt.Println(err)
			return -1
		}
		var sum float64 = 0
		var count = 0
		for _, grade := range enrollment.Grades {
			sum += float64(grade.Grade)
			count += 1
		}
		var average float64
		average = sum / float64(count)
		total += average
		totalCount += 1
	}
	if totalCount == 0 {
		return -1
	}
	return total / float64(totalCount)
}

func GetBestTeacherResults() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var courses []models.Course
		courses = GetAllCoursesForStatistics()
		var bestAverage float64 = -1
		var bestTeacher models.User
		for _, course := range courses {
			var averageGrade = GetAverageGradeByCourseID(c, course.ID)
			if averageGrade > bestAverage {
				bestAverage = averageGrade
				bestTeacher = *course.Proposer
			}
		}
		if bestAverage == -1 {
			c.JSON(http.StatusOK, "No grades available to determine best teacher")
		} else {
			c.JSON(http.StatusOK, gin.H{"bestTeacher": bestTeacher, "averageGrade": bestAverage})
		}
	}
}

func GetWorstTeacherResults() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var courses []models.Course
		courses = GetAllCoursesForStatistics()
		var worstAverage float64 = 10000
		var worstTeacher models.User
		for _, course := range courses {
			var averageGrade = GetAverageGradeByCourseID(c, course.ID)
			if averageGrade < worstAverage && averageGrade != -1 {
				worstAverage = averageGrade
				worstTeacher = *course.Proposer
			}
		}
		if worstAverage == -1 {
			c.JSON(http.StatusOK, "No grades available to determine worst teacher")
		} else {
			c.JSON(http.StatusOK, gin.H{"bestTeacher": worstTeacher, "averageGrade": worstAverage})
		}
	}
}
