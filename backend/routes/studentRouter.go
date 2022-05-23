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
	incomingRoutes.GET("/enrollments/grades/:courseid", controllers.ViewGrades())
	incomingRoutes.POST("/enrollments/signcontract/:academic_year_id", controllers.SignContract())
}
