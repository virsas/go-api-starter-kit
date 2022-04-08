package test

import (
	"database/sql"
	"go-api-starter-kit/utils/audit"
	"go-api-starter-kit/utils/db"
	"go-api-starter-kit/utils/logger"
	"go-api-starter-kit/utils/router"

	"github.com/gin-gonic/gin"
)

type TestAPI struct {
	Logger logger.LoggerHandler
	DB     *sql.DB
	Router *gin.Engine
	Audit  audit.AuditHandler
}

func NewTestAPI() (*TestAPI, error) {
	var err error
	env := &TestAPI{}

	l, err := logger.New()
	if err != nil {
		return nil, err
	}
	env.Logger = l

	d, err := db.New()
	if err != nil {
		return nil, err
	}
	if err := db.Migrate(d, "file://../../migrations"); err != nil {
		return nil, err
	}
	env.DB = d

	r, err := router.New("../../assets")
	if err != nil {
		return nil, err
	}
	env.Router = r

	a, err := NewTestAudit()
	if err != nil {
		return nil, err
	}
	env.Audit = a

	return env, nil
}
