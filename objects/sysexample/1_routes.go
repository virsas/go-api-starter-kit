package sysexample

import (
	"database/sql"
	"go-api-starter-kit/middlewares"
	"go-api-starter-kit/utils/logger"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, apiPath string, db *sql.DB, log logger.LoggerHandler) {
	ctrl := newController(db, log)

	r.GET(
		apiPath+"/sys"+"/examples/",
		middlewares.AllowAdmin(log),
		ctrl.list,
	)
	r.POST(
		apiPath+"/sys"+"/examples/",
		middlewares.AllowAdmin(log),
		validateExample(log),
		ctrl.create,
	)
	r.GET(
		apiPath+"/sys"+"/examples/:ID",
		middlewares.AllowAdmin(log),
		ctrl.show,
	)
	r.PATCH(
		apiPath+"/sys"+"/examples/:ID",
		middlewares.AllowAdmin(log),
		validateExample(log),
		ctrl.update,
	)
	r.DELETE(
		apiPath+"/sys"+"/examples/:ID",
		middlewares.AllowAdmin(log),
		ctrl.delete,
	)
}
