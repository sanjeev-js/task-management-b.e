package main

import (
	"task-management/config"
	"task-management/middlewares"
	"task-management/migrations"
	"task-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	migrations.Migrate()

	r := gin.Default()
	r.Use(middlewares.CorsConfig())

	routes.RegisterRoutes(r)

	r.POST("/login", loginHandler)

	r.Run(":8080")
}

//Checking if the route is working
func loginHandler(c *gin.Context){
	c.JSON(200,gin.H{"message":"Successfully logged in"})

}
