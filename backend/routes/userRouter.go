package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/courses", controllers.GetCourses())
	incomingRoutes.GET("/courses/:academic_year_id", controllers.GetCoursesByAcademicYear())
	incomingRoutes.GET("/proposed_courses", controllers.GetProposedCourses())
	incomingRoutes.GET("/groups/:academic_year_id", controllers.GetGroupsByAcademicYear())
	incomingRoutes.GET("/proposed_courses/:academic_year_id", controllers.GetProposedCoursesByAcademicYear())
	incomingRoutes.GET("/enrollments/getbyusername/:username", controllers.GetOptionalEnrollmentsByStudentUsername())
}
