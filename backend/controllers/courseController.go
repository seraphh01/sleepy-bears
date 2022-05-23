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
)

var courseCollection = database.OpenCollection(database.Client, "Course")
var proposedCourseCollection = database.OpenCollection(database.Client, "ProposedCourse")
var academicYearCollection = database.OpenCollection(database.Client, "AcademicYears")

func ApproveCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "CHIEF"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var course models.Course
		var amount models.Amount
		if err := c.BindJSON(&amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		courseID := c.Param("id")
		realCourseID, _ := primitive.ObjectIDFromHex(courseID)
		var studentEnrollmentCount = GetEnrollmentsCountByCourseID(c, realCourseID)

		if studentEnrollmentCount < 0 {
			return
		}

		if studentEnrollmentCount < 20 {
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
		ctype := "OPTIONAL"
		if *course.CourseType == ctype {
			course.MaxAmount = &amount
		} else {
			var max models.Amount
			max.Max = 2147483647
			course.MaxAmount = &max
		}

		validationErr := validate.Struct(course)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		err = DeleteEnrollmentsByCourseID(realCourseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func getCurrentAcademicYear() models.AcademicYear {
	currentTime := time.Now()
	oct, err := time.Parse(time.RFC822, "01 Oct "+strconv.Itoa(currentTime.Year())+" 0:00 UTC")
	if err != nil {
		fmt.Println("Invalid current time")
		return models.AcademicYear{}
	}
	var academicYear models.AcademicYear
	if currentTime.After(oct) {
		academicYear.StartDate = oct
		endDate, err := time.Parse(time.RFC822, "01 Jul "+strconv.Itoa(currentTime.Year()+1)+" 0:00 UTC")
		if err != nil {
			fmt.Println("Invalid end date")
			return models.AcademicYear{}
		}
		academicYear.EndDate = endDate
	} else {
		prevDate, err := time.Parse(time.RFC822, "01 Oct "+strconv.Itoa(currentTime.Year()-1)+" 0:00 UTC")
		if err != nil {
			fmt.Println("Invalid start date")
			return models.AcademicYear{}
		}
		academicYear.StartDate = prevDate
		endDate, err := time.Parse(time.RFC822, "01 Jul "+strconv.Itoa(currentTime.Year())+" 0:00 UTC")
		if err != nil {
			fmt.Println("Invalid end date")
			return models.AcademicYear{}
		}
		academicYear.EndDate = endDate
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	count, err := academicYearCollection.CountDocuments(ctx, bson.M{"start_date": academicYear.StartDate, "end_date": academicYear.EndDate})
	if err != nil {
		return models.AcademicYear{}
	}
	if count == 0 {
		academicYearCollection.InsertOne(ctx, academicYear)
	}
	return academicYear
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

		ctype := "OPTIONAL"
		course.CourseType = &ctype
		course.Proposer = &user
		if err := c.BindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		academicYear := getCurrentAcademicYear()
		if academicYear == (models.AcademicYear{}) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for current academic year"})
			return
		}
		course.AcademicYear = &academicYear

		count, err := proposedCourseCollection.CountDocuments(ctx, bson.M{"proposer._id": realUserID, "academic_year._id": course.AcademicYear.ID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count >= 2 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This teacher has already proposed two courses!"})
			return
		}

		countAccepted, err := courseCollection.CountDocuments(ctx, bson.M{"proposer._id": realUserID, "academic_year._id": course.AcademicYear.ID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if countAccepted+count >= 2 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This teacher has already proposed two courses!"})
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

func AddMandatoryCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "CHIEF"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var course models.Course
		var user models.User

		username := c.Param("teacher")
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctype := "MANDATORY"
		course.CourseType = &ctype
		course.Proposer = &user
		var max models.Amount
		max.Max = 2147483647
		course.MaxAmount = &max
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

		resultInsertionNumber, insertErr := courseCollection.InsertOne(ctx, course)
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

func GetCoursesByAcademicYear() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		yearId := c.Param("academic_year_id")
		realYearId, err := primitive.ObjectIDFromHex(yearId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var courses []models.Course
		cursor, err := courseCollection.Find(ctx, bson.M{"academic_year._id": realYearId})
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

func GetProposedCoursesByAcademicYear() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		yearId := c.Param("academic_year_id")
		realYearId, err := primitive.ObjectIDFromHex(yearId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var courses []models.Course
		cursor, err := proposedCourseCollection.Find(ctx, bson.M{"academic_year._id": realYearId})
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
			c.JSON(http.StatusOK, "No proposed courses available in this academic year")
		}
	}
}

func GetProposedCoursesByTeacherUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "TEACHER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		username := c.Param("username")
		var courses []models.Course
		cursor, err := proposedCourseCollection.Find(ctx, bson.M{"proposer.username": username})
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
			c.JSON(http.StatusOK, "No proposed courses available for this teacher")
		}
	}
}

func GetAllCoursesForStatistics() []models.Course {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var courses []models.Course
	cursor, err := courseCollection.Find(ctx, bson.M{})
	if err != nil {
		return []models.Course{}
	}
	for cursor.Next(ctx) {
		var course models.Course
		err := cursor.Decode(&course)
		if err != nil {
			fmt.Println("error for some reason")
			return []models.Course{}
		}
		courses = append(courses, course)
	}
	if len(courses) > 0 {
		return courses
	} else {
		return []models.Course{}
	}
}

func GetCoursesByYearForStatistics(year int) []models.Course {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var courses []models.Course
	cursor, err := courseCollection.Find(ctx, bson.M{"year_of_study": year})
	if err != nil {
		return []models.Course{}
	}
	for cursor.Next(ctx) {
		var course models.Course
		err := cursor.Decode(&course)
		if err != nil {
			fmt.Println(err)
			return []models.Course{}
		}
		courses = append(courses, course)
	}
	if len(courses) > 0 {
		return courses
	} else {
		return []models.Course{}
	}
}

func GetCoursesBySemesterForStatistics(semester int) []models.Course {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var courses []models.Course
	cursor, err := courseCollection.Find(ctx, bson.M{"semester": semester})
	if err != nil {
		return []models.Course{}
	}
	for cursor.Next(ctx) {
		var course models.Course
		err := cursor.Decode(&course)
		if err != nil {
			fmt.Println(err)
			return []models.Course{}
		}
		courses = append(courses, course)
	}
	if len(courses) > 0 {
		return courses
	} else {
		return []models.Course{}
	}
}

func GetStudentsByCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var students []models.User
		var course models.Course
		courseId := c.Param("courseid")
		realCourseId, _ := primitive.ObjectIDFromHex(courseId)
		err := courseCollection.FindOne(ctx, bson.M{"_id": realCourseId}).Decode(&course)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		students = GetStudentsFromCourses([]models.Course{course})
		if len(students) > 0 {
			c.JSON(http.StatusOK, students)
		} else {
			c.JSON(http.StatusOK, "No students enrolled in this course")
		}
	}
}
