package health

import (
	"go-api-starter-kit/config"

	"github.com/gin-gonic/gin"
)

func (env *env) show(ctx *gin.Context) {
	s := &service{db: env.db, log: env.log}
	dbstatus := s.show()
	ctx.JSON(config.OK_STATUS, gin.H{
		"server":   "OK",
		"database": dbstatus,
	})
}
