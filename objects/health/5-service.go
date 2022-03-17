package health

import (
	"database/sql"

	"go.uber.org/zap"
)

type service struct {
	db  *sql.DB
	log *zap.Logger
}

func (s *service) show() string {
	m := &model{db: s.db, log: s.log}
	dbstatus := m.show()

	return dbstatus
}
