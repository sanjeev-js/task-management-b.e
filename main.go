package main

import (
	"net/http"
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

	r.POST("/login", staticCheckUser)

	r.Run(":8080")
}

//Checking if the route is working
var staticUser = map[string]struct {
	Password string
	Role     string
}{
	"admin@gmail.com": {Password: "admin@123", Role: "admin"},
	"user@gmail.com":  {Password: "user@123", Role: "user"},
}

func staticCheckUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//Bind Json Input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Validate Credentials

	user, exists := staticUser[input.Email]
	if !exists || user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"role": user.Role,
	})

}
