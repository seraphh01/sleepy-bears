package main

import (
	"backend/routes"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefaultConfig() cors.Config {
	conf := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	conf.AllowAllOrigins = true
	return conf
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.New(DefaultConfig()))

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.StudentRoutes(router)
	routes.TeacherRoutes(router)
	routes.ChiefRoutes(router)
	routes.AdminRoutes(router)

	router.Run("127.0.0.1:" + port)
}
