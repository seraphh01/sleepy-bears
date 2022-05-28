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
	incomingRoutes.GET("/users/group/:group_id", controllers.GetStudentsByGroup())
	incomingRoutes.POST("/users/generate/:type/:groupid", controllers.GenerateUsers())
	incomingRoutes.PUT("/users/update/:username", controllers.UpdateUser())
	incomingRoutes.DELETE("/users/remove/:username", controllers.DeleteUser())
	incomingRoutes.POST("/groups/add", controllers.AddGroup())
	incomingRoutes.POST("/groups/add_student/:groupnumber/:username", controllers.AddStudentToGroup())
	incomingRoutes.GET("/users/studentsbygroup/performancedesc/:groupid", controllers.GetAllStudentsFromGroupSortedByAverageGradeDesc())

	//this one works, but it would be better if you simply called the previous one for each group
	incomingRoutes.GET("/users/allgroups/allstudents/performancedesc", controllers.AllStudentsFromAllGroupsSortedByPerformanceDesc())

	incomingRoutes.GET("/users/studentsbyyear/performancedesc/:year", controllers.GetAllStudentsFromYearSortedByAverageGradeDesc())
	incomingRoutes.GET("/users/studentsbysemester/performancedesc/:semester", controllers.GetAllStudentsFromSemesterSortedByAverageGradeDesc())

	incomingRoutes.POST("/makechief/:username", controllers.MakeTeacherChief())
}
