package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func ChiefRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/proposed_courses", controllers.GetProposedCourses())
	incomingRoutes.POST("/courses/approve/:id", controllers.AddCourse())
}
