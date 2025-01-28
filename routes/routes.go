package routes

import (
	"task-management/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		// authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", controllers.CreateUser)
	}

	// protectedGroup := router.Group("/user")
	// protectedGroup.Use(middlewares.AuthMiddleware())
	// {
	// 	protectedGroup.GET("/profile", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{"message": "Welcome to the protected route!"})
	// 	})
	// }
}
