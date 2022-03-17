package health

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routes(r *gin.Engine, apiPath string, db *sql.DB, log *zap.Logger) {
	health := &env{db: db, log: log}

	r.GET(
		apiPath+"/health",
		health.show,
	)
}
