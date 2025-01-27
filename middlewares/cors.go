package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsConfig() gin.HandlerFunc {
//Allowing Cross Origin access for the frontend	
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "UPDATE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorisation"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

}
