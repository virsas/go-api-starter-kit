package example

import (
	"database/sql"
	"go-api-starter-kit/middlewares"
	"go-api-starter-kit/utils/logger"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, apiPath string, db *sql.DB, log logger.LoggerHandler) {
	ctrl := newController(db, log)

	r.GET(
		apiPath+"/examples/",
		middlewares.AllowRoles(log, []string{"admin"}),
		ctrl.list,
	)
	r.POST(
		apiPath+"/examples/",
		middlewares.AllowRoles(log, []string{"admin"}),
		validateExample(log),
		ctrl.create,
	)
	r.GET(
		apiPath+"/examples/:ID",
		middlewares.AllowRoles(log, []string{"admin"}),
		ctrl.show,
	)
	r.PATCH(
		apiPath+"/examples/:ID",
		middlewares.AllowRoles(log, []string{"admin"}),
		validateExample(log),
		ctrl.update,
	)
	r.DELETE(
		apiPath+"/examples/:ID",
		middlewares.AllowRoles(log, []string{"admin"}),
		ctrl.delete,
	)
}
