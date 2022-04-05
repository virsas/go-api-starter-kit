package health

import (
	"database/sql"
	"go-api-starter-kit/config"
	"go-api-starter-kit/utils/logger"

	"github.com/gin-gonic/gin"
)

type controller struct {
	db  *sql.DB
	log *logger.Logger
	s   *service
}

func newController(db *sql.DB, log *logger.Logger) *controller {
	s := newService(db, log)
	return &controller{db: db, log: log, s: s}
}

func (ctrl *controller) show(ctx *gin.Context) {
	dbstatus := ctrl.s.show()
	ctx.JSON(config.OK_STATUS, gin.H{
		"server":   "OK",
		"database": dbstatus,
	})
}
