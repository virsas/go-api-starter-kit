package router

import (
	"os"
	"strconv"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func New() (*gin.Engine, error) {
	logger, _ := zap.NewProduction()

	var debugLogging bool = false
	debugLoggingValue, debugLoggingPresent := os.LookupEnv("ROUTER_DEBUG")
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

	r.Use(getCors())
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	return r, nil
}
