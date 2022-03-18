package health

import (
	"database/sql"

	"go.uber.org/zap"
)

type model struct {
	db  *sql.DB
	log *zap.Logger
}

func newModel(db *sql.DB, log *zap.Logger) *model {
	return &model{db: db, log: log}
}

func (m *model) show() string {
	err := m.db.Ping()
	if err != nil {
		return "Not OK"
	}
	return "OK"
}
