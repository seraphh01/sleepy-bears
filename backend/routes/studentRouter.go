package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.POST("/enrollments/add/:courseid", controllers.AddEnrollment())
	incomingRoutes.GET("/enrollments/grades/:courseid", controllers.ViewGrades())
	incomingRoutes.POST("/enrollments/signcontract/:year", controllers.SignContract())
}
