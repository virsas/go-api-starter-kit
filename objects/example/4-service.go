package example

import (
	"context"
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

func (s *service) list() ([]examplemodel, error) {
	var examples []examplemodel = []examplemodel{}
	examples, err := s.m.list()
	if err != nil {
		return examples, err
	}
	return examples, nil
}

func (s *service) create(c context.Context, example examplereq) error {
	err := s.m.create(c, example)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) show(id int64) (examplemodel, error) {
	var example examplemodel
	example, err := s.m.show(id)
	if err != nil {
		return example, err
	}

	return example, nil
}

func (s *service) update(c context.Context, id int64, example examplereq) error {
	err := s.m.update(c, id, example)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) delete(c context.Context, id int64) error {
	err := s.m.delete(c, id)
	if err != nil {
		return err
	}

	return nil
}
