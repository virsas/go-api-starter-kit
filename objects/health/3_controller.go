package health

import (
	"database/sql"
	"go-api-starter-kit/utils/logger"
	"go-api-starter-kit/utils/vars"

	"github.com/gin-gonic/gin"
)

type controller struct {
	db  *sql.DB
	log logger.LoggerHandler
	m   *model
}

func newController(db *sql.DB, log logger.LoggerHandler) *controller {
	m := newModel(db, log)
	return &controller{db: db, log: log, m: m}
}

func (ctrl *controller) show(ctx *gin.Context) {
	dbstatus := ctrl.m.show()
	ctx.JSON(vars.STATUS_OK_CODE, gin.H{
		"server":   "OK",
		"database": dbstatus,
	})
}
