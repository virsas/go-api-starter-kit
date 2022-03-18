package health

import (
	"database/sql"
	"go-api-starter-kit/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type controller struct {
	db  *sql.DB
	log *zap.Logger
	s   *service
}

func newController(db *sql.DB, log *zap.Logger) *controller {
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
