package router

import (
	"context"
	"database/sql"

	"go.uber.org/zap"
)

type Env struct {
	db  *sql.DB
	log *zap.Logger
	ctx context.Context
}
