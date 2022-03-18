package example

import (
	"context"
	"database/sql"

	"go.uber.org/zap"
)

type service struct {
	db  *sql.DB
	log *zap.Logger
	ctx context.Context
}

func (s *service) list() ([]examplemodel, error) {
	var examples []examplemodel = []examplemodel{}
	m := &model{db: s.db, log: s.log, ctx: s.ctx}
	examples, err := m.list()
	if err != nil {
		return examples, err
	}
	return examples, nil
}

func (s *service) create(example examplereq) error {
	m := &model{db: s.db, log: s.log, ctx: s.ctx}
	err := m.create(example)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) show(id int64) (examplemodel, error) {
	var example examplemodel
	m := &model{db: s.db, log: s.log, ctx: s.ctx}
	example, err := m.show(id)
	if err != nil {
		return example, err
	}

	return example, nil
}

func (s *service) update(id int64, example examplereq) error {
	m := &model{db: s.db, log: s.log, ctx: s.ctx}
	err := m.update(id, example)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) delete(id int64) error {
	m := &model{db: s.db, log: s.log, ctx: s.ctx}
	err := m.delete(id)
	if err != nil {
		return err
	}

	return nil
}
