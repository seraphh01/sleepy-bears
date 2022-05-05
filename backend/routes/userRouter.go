package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.POST("/users/generate/:type", controllers.GenerateUsers())
	incomingRoutes.GET("/users/:type", controllers.GetUsers())
	incomingRoutes.GET("/user/:username", controllers.GetUser())
	incomingRoutes.PUT("/users/update/:username", controllers.UpdateUser())
	incomingRoutes.DELETE("/users/remove/:username", controllers.DeleteUser())
}
