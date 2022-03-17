package example

import (
	"context"
	"database/sql"

	"go.uber.org/zap"
)

// environment configuration
type env struct {
	db  *sql.DB
	log *zap.Logger
	ctx context.Context
}
