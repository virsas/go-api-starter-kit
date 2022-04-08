package health

import (
	"database/sql"
	"go-api-starter-kit/utils/logger"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, apiPath string, db *sql.DB, log logger.LoggerHandler) {
	ctrl := newController(db, log)

	r.GET(
		"/health",
		ctrl.show,
	)
	r.GET(
		apiPath+"/v1"+"/status",
		ctrl.show,
	)
}
