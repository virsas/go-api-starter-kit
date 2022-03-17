package health

import (
	"database/sql"

	"go.uber.org/zap"
)

// environment configuration
type env struct {
	db  *sql.DB
	log *zap.Logger
}
