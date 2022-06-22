package router

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getCors() gin.HandlerFunc {

	var corsAllowedOrigin string = "https://*"
	corsAllowedOriginValue, corsAllowedOriginPresent := os.LookupEnv("CORS_ALLOW_ORIGIN")
	if corsAllowedOriginPresent {
		corsAllowedOrigin = corsAllowedOriginValue
	}

	c := cors.New(cors.Config{
		AllowOrigins:     []string{corsAllowedOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "APIKey", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	return c

}
