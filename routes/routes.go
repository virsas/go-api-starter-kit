package routes

import (
	"database/sql"
	"go-api-starter-kit/middlewares"
	"go-api-starter-kit/objects/example"
	"go-api-starter-kit/objects/health"
	"go-api-starter-kit/objects/sysexample"
	"go-api-starter-kit/utils/audit"
	"go-api-starter-kit/utils/logger"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func AddRoutes(r *gin.Engine, d *sql.DB, l logger.LoggerHandler, a audit.AuditHandler) {
	var apiPath string = ""
	apiPathValue, apiPathPresent := os.LookupEnv("API_PATH")
	if apiPathPresent {
		apiPath = apiPathValue
	}

	// order is important
	r.Use(favicon.New("assets/favicon.ico"))
	health.Routes(r, apiPath, d, l)

	// all routes should be below authentication, only health route is available without authentication
	r.Use(middlewares.Auth("./keys", "test", l))
	r.Use(middlewares.User(d, l))
	r.Use(middlewares.Log(a, l))

	example.Routes(r, apiPath, d, l)
	sysexample.Routes(r, apiPath, d, l)
}
