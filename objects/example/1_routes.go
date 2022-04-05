package example

import (
	"database/sql"
	"go-api-starter-kit/middlewares"
	"go-api-starter-kit/utils/logger"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, apiPath string, db *sql.DB, log *logger.Logger) {
	ctrl := newController(db, log)

	r.GET(
		"/examples/",
		middlewares.AllowRoles(log, []string{"admin"}),
		ctrl.list,
	)
	r.POST(
		"/examples/",
		middlewares.AllowRoles(log, []string{"admin"}),
		validateExample(log),
		ctrl.create,
	)
	r.GET(
		"/examples/:ID",
		middlewares.AllowRoles(log, []string{"admin"}),
		ctrl.show,
	)
	r.PATCH(
		"/examples/:ID",
		middlewares.AllowRoles(log, []string{"admin"}),
		validateExample(log),
		ctrl.update,
	)
	r.DELETE(
		"/examples/:ID",
		middlewares.AllowRoles(log, []string{"admin"}),
		ctrl.delete,
	)
}
