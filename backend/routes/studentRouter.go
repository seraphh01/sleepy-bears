package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.POST("/enrollments/add/:courseid", controllers.AddEnrollment())
	incomingRoutes.DELETE("/enrollments/remove/:courseid", controllers.RemoveEnrollment())
	incomingRoutes.GET("/enrollments/grades/:courseid", controllers.ViewGradesByCourse())
	incomingRoutes.GET("/enrollments/grades", controllers.ViewAllGrades())
	incomingRoutes.POST("/enrollments/signcontract/:academic_year_id", controllers.SignContract())
}
