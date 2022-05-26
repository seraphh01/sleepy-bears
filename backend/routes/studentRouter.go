package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.POST("/enrollments/add/:courseid", controllers.AddEnrollment())
	incomingRoutes.POST("/enrollments/add/by-year/:year_of_study", controllers.AddEnrollmentsToYearOfStudy())
	incomingRoutes.DELETE("/enrollments/remove/:courseid", controllers.RemoveEnrollment())
	incomingRoutes.GET("/enrollments/grades/:courseid", controllers.ViewGradesByCourse())
	incomingRoutes.GET("/enrollments/grades", controllers.ViewAllGrades())
	incomingRoutes.POST("/enrollments/signcontract/:academic_year_id", controllers.SignContract())
	incomingRoutes.GET("/enrollments/grades/by-year/:username", controllers.ViewAllGradesByYear())
}
