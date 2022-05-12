package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func ChiefRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.POST("/courses/approve/:id", controllers.ApproveCourse())
	incomingRoutes.POST("/courses/addmandatory/:teacher", controllers.AddMandatoryCourse())
}
