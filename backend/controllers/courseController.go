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
		course.ID = realCourseID

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
	//create a primitive datetime given a date

	var july = time.Date(currentTime.Year(), time.July, 1, 0, 0, 0, 0, time.UTC)
	var oct = time.Date(currentTime.Year(), time.October, 1, 0, 0, 0, 0, time.UTC)
	var academicYear models.AcademicYear
	if currentTime.After(july) {
		academicYear.StartDate = primitive.NewDateTimeFromTime(oct)
		endDate := time.Date(currentTime.Year()+1, time.July, 1, 0, 0, 0, 0, time.UTC)
		academicYear.EndDate = primitive.NewDateTimeFromTime(endDate)
	} else {
		prevDate := time.Date(currentTime.Year()-1, time.October, 1, 0, 0, 0, 0, time.UTC)
		academicYear.StartDate = primitive.NewDateTimeFromTime(prevDate)
		academicYear.EndDate = primitive.NewDateTimeFromTime(july)
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var foundAcademicYear models.AcademicYear
	err := academicYearCollection.FindOne(ctx, bson.M{"startdate": academicYear.StartDate, "enddate": academicYear.EndDate}).Decode(&foundAcademicYear)
	if err != nil {
		academicYear.ID = primitive.NewObjectID()
		_, err := academicYearCollection.InsertOne(ctx, academicYear)
		if err != nil {
			return models.AcademicYear{}
		}
		return academicYear
	} else {
		return foundAcademicYear
	}
}

func GetCurrentAcademicYear() gin.HandlerFunc {
	return func(c *gin.Context) {
		academicYear := getCurrentAcademicYear()
		if academicYear == (models.AcademicYear{}) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for current academic year"})
			return
		}
		c.JSON(http.StatusOK, academicYear)
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

		username := c.Param("username")
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		realUserID, _ := primitive.ObjectIDFromHex(user.ID.Hex())
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

		count, err := proposedCourseCollection.CountDocuments(ctx, bson.M{"proposer._id": realUserID, "academicyear._id": course.AcademicYear.ID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count >= 2 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This teacher has already proposed two courses!"})
			return
		}

		countAccepted, err := courseCollection.CountDocuments(ctx, bson.M{"proposer._id": realUserID, "academicyear._id": course.AcademicYear.ID})
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

		academicYear := getCurrentAcademicYear()
		course.AcademicYear = &academicYear

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
		year_of_study, err := strconv.Atoi(c.Param("year_of_study"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var courses []models.Course
		cursor, err := courseCollection.Find(ctx, bson.M{"academicyear._id": realYearId, "yearofstudy": year_of_study})
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
			c.JSON(http.StatusOK, "No courses available in this academic year and year of study")
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
		year_of_study, err := strconv.Atoi(c.Param("year_of_study"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var courses []models.Course
		cursor, err := proposedCourseCollection.Find(ctx, bson.M{"academicyear._id": realYearId, "yearofstudy": year_of_study})
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
			c.JSON(http.StatusOK, "No proposed courses available in this academic year and year of study")
		}
	}
}

func GetCoursesByYearOfStudy() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		year_of_study, err := strconv.Atoi(c.Param("year_of_study"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var courses []models.Course
		cursor, err := courseCollection.Find(ctx, bson.M{"yearofstudy": year_of_study})
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
			c.JSON(http.StatusOK, "No courses available in this year of study")
		}
	}
}

func GetProposedCoursesByYearOfStudy() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		year_of_study, err := strconv.Atoi(c.Param("year_of_study"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var courses []models.Course
		cursor, err := proposedCourseCollection.Find(ctx, bson.M{"yearofstudy": year_of_study})
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
			c.JSON(http.StatusOK, "No proposed courses available in this year of study")
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

func GetMandatoryCoursesByTeacherUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "TEACHER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		username := c.Param("username")
		var courses []models.Course
		cursor, err := courseCollection.Find(ctx, bson.M{"proposer.username": username})
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
	cursor, err := courseCollection.Find(ctx, bson.M{"yearofstudy": year})
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
			err2 := proposedCourseCollection.FindOne(ctx, bson.M{"_id": realCourseId}).Decode(&course)
			if err2 != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		students = GetStudentsFromCourses([]models.Course{course})
		if len(students) > 0 {
			c.JSON(http.StatusOK, students)
		} else {
			c.JSON(http.StatusOK, "No students enrolled in this course")
		}
	}
}

func GetAverageGradeAtCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "TEACHER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		courseid := c.Param("courseid")
		realCourseId, _ := primitive.ObjectIDFromHex(courseid)
		cursor, err := enrollmentCollection.Find(ctx, bson.M{"course._id": realCourseId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var total float64 = 0
		var totalCount = 0
		for cursor.Next(ctx) {
			var enrollment models.Enrollment
			err := cursor.Decode(&enrollment)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No grades available"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"average": total / float64(totalCount)})
	}
}
