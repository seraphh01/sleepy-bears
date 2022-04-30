package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func CourseRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/courses", controllers.GetCourses())
	incomingRoutes.GET("/courses/:year", controllers.GetCoursesByYear())
	incomingRoutes.GET("/proposed_courses", controllers.GetProposedCourses())
	incomingRoutes.POST("/courses/approve/:id", controllers.AddCourse())
	incomingRoutes.POST("/proposed_courses/add", controllers.AddProposedCourse())
}
