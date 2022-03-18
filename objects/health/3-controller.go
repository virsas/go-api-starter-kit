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
}

func (ctrl *controller) show(ctx *gin.Context) {
	s := &service{db: ctrl.db, log: ctrl.log}
	dbstatus := s.show()
	ctx.JSON(config.OK_STATUS, gin.H{
		"server":   "OK",
		"database": dbstatus,
	})
}
