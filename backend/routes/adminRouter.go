package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.POST("/users/generate/:type", controllers.GenerateUsers())
	incomingRoutes.GET("/users/:type", controllers.GetUsers())
	incomingRoutes.GET("/user/:username", controllers.GetUser())
	incomingRoutes.PUT("/user/update/:username", controllers.UpdateUser())
	incomingRoutes.DELETE("/user/remove/:username", controllers.DeleteUser())
	incomingRoutes.POST("/groups/add", controllers.AddGroup())
}
