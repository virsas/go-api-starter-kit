package routes

import (
	"context"
	"database/sql"
	"go-api-starter-kit/middlewares"
	"go-api-starter-kit/objects/example"
	"go-api-starter-kit/objects/health"
	"go-api-starter-kit/utils"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddRoutes(r *gin.Engine, db *sql.DB, logger *zap.Logger, ctx context.Context, audit *utils.Audit) {
	var apiPath string = ""
	apiPathValue, apiPathPresent := os.LookupEnv("API_PATH")
	if apiPathPresent {
		apiPath = apiPathValue
	}

	// order is important
	health.Routes(r, apiPath, db, logger)

	// all routes should be below authentication, only health route is available without authentication
	r.Use(middlewares.Auth(true, logger))
	r.Use(middlewares.User(db, logger))
	r.Use(middlewares.Log(audit, logger))

	example.Routes(r, apiPath, db, logger, ctx)
}
