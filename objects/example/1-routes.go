package example

import (
	"context"
	"database/sql"
	"go-api-starter-kit/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routes(r *gin.Engine, apiPath string, db *sql.DB, log *zap.Logger, ctx context.Context) {
	ctrl := &controller{db: db, log: log, ctx: ctx}

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
