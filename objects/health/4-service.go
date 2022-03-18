package health

import (
	"database/sql"

	"go.uber.org/zap"
)

type service struct {
	db  *sql.DB
	log *zap.Logger
	m   *model
}

func newService(db *sql.DB, log *zap.Logger) *service {
	m := newModel(db, log)
	return &service{db: db, log: log, m: m}
}

func (s *service) show() string {
	dbstatus := s.m.show()

	return dbstatus
}
