package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"go.uber.org/zap"
)

func InitRouter() (*gin.Engine, error) {
	logger, _ := zap.NewProduction()

	// Enable debugging based on ENV configuration, disabled by default
	var debugLogging bool = false
	debugLoggingValue, debugLoggingPresent := os.LookupEnv("GIN_DEBUG")
	if debugLoggingPresent {
		debugLoggingValueBool, err := strconv.ParseBool(debugLoggingValue)
		if err != nil {
			debugLogging = false
		}
		debugLogging = debugLoggingValueBool
	}

	if debugLogging {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	r := gin.Default()

	cors := cors.New(cors.Config{
		AllowOrigins:     []string{"https://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "APIKey", "Content-Type", "X-CSRF-Token", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	r.Use(cors)
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(favicon.New("./assets/favicon.ico"))
	return r, nil
}
