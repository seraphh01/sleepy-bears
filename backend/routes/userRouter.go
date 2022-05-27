package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/courses", controllers.GetCourses())
	incomingRoutes.GET("/courses/academic_year/:academic_year_id/:year_of_study", controllers.GetCoursesByAcademicYear())
	incomingRoutes.GET("/courses/year_of_study/:year_of_study", controllers.GetCoursesByYearOfStudy())
	incomingRoutes.GET("/proposed_courses", controllers.GetProposedCourses())
	incomingRoutes.GET("/groups/:academic_year_id", controllers.GetGroupsByAcademicYear())
	incomingRoutes.GET("/proposed_courses/academic_year/:academic_year_id/:year_of_study", controllers.GetProposedCoursesByAcademicYear())
	incomingRoutes.GET("/proposed_courses/year_of_study/:year_of_study", controllers.GetProposedCoursesByYearOfStudy())
	incomingRoutes.GET("/enrollments/getbyusername/:username", controllers.GetOptionalEnrollmentsByStudentUsername())
	incomingRoutes.GET("/enrollments/getmandatorybyusername/:username/:yearofstudy", controllers.GetMandatoryEnrollmentsByStudentUsername())
	incomingRoutes.GET("/academic_year", controllers.GetCurrentAcademicYear())
}
