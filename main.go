package main

import (
	"task-management/config"
	"task-management/migrations"
	"task-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	migrations.Migrate()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
