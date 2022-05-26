package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func TeacherRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.POST("/proposed_courses/add/:username", controllers.AddProposedCourse())
	incomingRoutes.POST("/grades/add/:studentusername/:courseid", controllers.GradeStudent())
	incomingRoutes.GET("/proposed_courses/getby/:username", controllers.GetProposedCoursesByTeacherUsername())
	incomingRoutes.GET("/students/bycourse/:courseid", controllers.GetStudentsByCourse())
	incomingRoutes.GET("/students/average/:courseid", controllers.GetAverageGradeAtCourse())
	incomingRoutes.GET("/students/grades/:courseid", controllers.ViewGradesByCourseAllStudents())
	incomingRoutes.GET("/mandatorycourses/getby/:username", controllers.GetMandatoryCoursesByTeacherUsername())
}
