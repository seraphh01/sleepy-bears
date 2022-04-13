package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func CourseRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/courses", controllers.GetCourses())
	incomingRoutes.POST("/courses/add", controllers.AddCourse())
	//incomingRoutes.DELETE("/users/remove/:username", controllers.DeleteUser())
}
