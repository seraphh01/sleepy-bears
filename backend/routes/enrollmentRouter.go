package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func EnrollmentRouter(incomingRoutes *gin.Engine) {
	//incomingRoutes.GET("/enrollments", controllers.GetEnrollments())
	incomingRoutes.POST("/enrollments/add_optional/:courseid", controllers.AddOptionalEnrollment())
}
