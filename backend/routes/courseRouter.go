package routes

import (
  "backend/controllers"
  "github.com/gin-gonic/gin"
)

func CourseRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.GET("/courses", controllers.GetCourses())
  incomingRoutes.GET("/proposed_courses", controllers.GetProposedCourses())
  incomingRoutes.POST("/courses/add", controllers.AddCourse())
  incomingRoutes.POST("/proposed_courses/add", controllers.AddProposedCourse())
  //incomingRoutes.DELETE("/users/remove/:username", controllers.DeleteUser())
}
