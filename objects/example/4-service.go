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

func (s *service) list(aid int) ([]Example, error) {
	var examples []Example = []Example{}
	examples, err := s.m.list(aid)
	if err != nil {
		return examples, err
	}
	return examples, nil
}

func (s *service) create(c context.Context, example ExampleInput, aid int) error {
	err := s.m.create(c, example, aid)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) show(id int64, aid int) (Example, error) {
	var example Example
	example, err := s.m.show(id, aid)
	if err != nil {
		return example, err
	}

	return example, nil
}

func (s *service) update(c context.Context, id int64, example ExampleInput, aid int) error {
	err := s.m.update(c, id, example, aid)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) delete(c context.Context, id int64, aid int) error {
	err := s.m.delete(c, id, aid)
	if err != nil {
		return err
	}

	return nil
}
