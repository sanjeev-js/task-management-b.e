package routes

import (
	"task-management/controllers"
	"task-management/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
	}

	protectedGroup := router.Group("/user")
	protectedGroup.Use(middlewares.AuthMiddleware())
	{
		protectedGroup.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the protected route!"})
		})
	}
}
