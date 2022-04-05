package health

import (
	"database/sql"
	"go-api-starter-kit/utils/logger"
)

type model struct {
	db  *sql.DB
	log *logger.Logger
}

func newModel(db *sql.DB, log *logger.Logger) *model {
	return &model{db: db, log: log}
}

func (m *model) show() string {
	err := m.db.Ping()
	if err != nil {
		return "Not OK"
	}
	return "OK"
}
