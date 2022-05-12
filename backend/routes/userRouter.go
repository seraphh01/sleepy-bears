package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/courses", controllers.GetCourses())
	incomingRoutes.GET("/courses/:year", controllers.GetCoursesByYear())
	incomingRoutes.GET("/proposed_courses", controllers.GetProposedCourses())
	incomingRoutes.GET("/groups/:year", controllers.GetGroupsByYear())
	incomingRoutes.GET("/proposed_courses/:year", controllers.GetProposedCoursesByYear())
}
