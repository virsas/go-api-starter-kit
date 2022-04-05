package health

import (
	"database/sql"
	"go-api-starter-kit/utils/logger"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, apiPath string, db *sql.DB, log *logger.Logger) {
	ctrl := newController(db, log)

	r.GET(
		apiPath+"/health",
		ctrl.show,
	)
}
