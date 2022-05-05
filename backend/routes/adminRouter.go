package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/users/type/:type", controllers.GetUsers())
	incomingRoutes.GET("/users/username/:username", controllers.GetUser())
	incomingRoutes.GET("/users/group/:groupnumber", controllers.GetStudentsByGroup())
	incomingRoutes.POST("/users/generate/:type", controllers.GenerateUsers())
	incomingRoutes.PUT("/users/update/:username", controllers.UpdateUser())
	incomingRoutes.DELETE("/users/remove/:username", controllers.DeleteUser())
	incomingRoutes.POST("/groups/add", controllers.AddGroup())
	incomingRoutes.POST("/groups/add_student/:groupnumber/:username", controllers.AddStudentToGroup())
}
