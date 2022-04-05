package health

import (
	"database/sql"
	"go-api-starter-kit/utils/logger"
)

type service struct {
	db  *sql.DB
	log *logger.Logger
	m   *model
}

func newService(db *sql.DB, log *logger.Logger) *service {
	m := newModel(db, log)
	return &service{db: db, log: log, m: m}
}

func (s *service) show() string {
	dbstatus := s.m.show()

	return dbstatus
}
