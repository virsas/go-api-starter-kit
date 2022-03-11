package router

import (
	"context"
	"database/sql"
	"go-api-starter-kit/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AddRoutes function
func AddRoutes(r *gin.Engine, db *sql.DB, logger *zap.Logger, ctx context.Context) {
	var apiPath string = ""
	apiPathValue, apiPathPresent := os.LookupEnv("API_PATH")
	if apiPathPresent {
		apiPath = apiPathValue
	}

	var verPath string = ""
	verPathValue, verPathPresent := os.LookupEnv("VER_PATH")
	if verPathPresent {
		verPath = verPathValue
	}

	// Health
	r.GET(apiPath+verPath+"/health", healthShow)

	env := &Env{db: db, log: logger, ctx: ctx}
	// Controllers
	r.GET(apiPath+verPath+"/examples", middlewares.Example(), env.exampleList)
	r.POST(apiPath+verPath+"/examples", middlewares.Example(), env.exampleCreate)
	r.GET(apiPath+verPath+"/examples/:ID", middlewares.Example(), env.exampleShow)
	r.PUT(apiPath+verPath+"/examples/:ID", middlewares.Example(), env.exampleUpdate)
	r.DELETE(apiPath+verPath+"/examples/:ID", middlewares.Example(), env.exampleDelete)
}
