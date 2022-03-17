package example

import (
	"context"
	"database/sql"
	"go-api-starter-kit/config"
	"time"

	"go.uber.org/zap"
)

type model struct {
	db  *sql.DB
	log *zap.Logger
	ctx context.Context
}

// examplereq struct
type examplereq struct {
	Name *string `json:"name" validate:"required,min=2,max=255"`
}

// examplemodel struct
type examplemodel struct {
	ID *int64 `json:"id"`
	examplereq
	Updatedat *time.Time `json:"updatedat"`
	Createdat *time.Time `json:"createdat"`
}

func (m *model) list() ([]examplemodel, error) {
	var example examplemodel
	var examples []examplemodel = []examplemodel{}

	examplesQuery, err := m.db.Query("SELECT id, name, updated_at, created_at FROM examples ORDER BY id DESC;")
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return nil, config.NewDBError(err)
	}
	defer examplesQuery.Close()

	for examplesQuery.Next() {
		err = examplesQuery.Scan(&example.ID, &example.Name, &example.Updatedat, &example.Createdat)
		if err != nil {
			return nil, config.NewDBError(err)
		}

		examples = append(examples, example)
	}

	return examples, nil
}

func (m *model) create(model examplereq) error {
	tx, err := m.db.BeginTx(m.ctx, nil)
	if err != nil {
		m.log.Error("apiIssue", zap.Error(err))
		return config.NewServerError(err)
	}

	_, err = tx.ExecContext(m.ctx, "INSERT INTO examples (name, updated_at, created_at) VALUES (?,?,?);", model.Name, time.Now(), time.Now())
	if err != nil {
		tx.Rollback()
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	err = tx.Commit()
	if err != nil {
		m.log.Error("apiIssue", zap.Error(err))
		return config.NewServerError(err)
	}

	return nil
}

func (m *model) show(id int64) (examplemodel, error) {
	var example examplemodel
	err := m.db.QueryRow("SELECT id, name, updated_at, created_at FROM examples WHERE id=?;", id).Scan(&example.ID, &example.Name, &example.Updatedat, &example.Createdat)
	if err != nil {
		if err == sql.ErrNoRows {
			m.log.Error("notFound", zap.Error(err))
			return example, config.NewNotfoundError(err)
		}
		m.log.Error("dbError", zap.Error(err))
		return example, config.NewDBError(err)
	}

	return example, nil
}

func (m *model) update(id int64, model examplereq) error {
	tx, err := m.db.BeginTx(m.ctx, nil)
	if err != nil {
		m.log.Error("apiIssue", zap.Error(err))
		return config.NewServerError(err)
	}

	var example examplemodel
	err = tx.QueryRowContext(m.ctx, "SELECT id FROM examples WHERE id=?;", id).Scan(&example.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			m.log.Error("notFound", zap.Error(err))
			return config.NewNotfoundError(err)
		}
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	_, err = tx.ExecContext(m.ctx, "UPDATE examples SET name=?, updated_at=? WHERE id=?;", model.Name, time.Now(), example.ID)
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	err = tx.Commit()
	if err != nil {
		m.log.Error("apiIssue", zap.Error(err))
		return config.NewServerError(err)
	}

	return nil
}

func (m *model) delete(id int64) error {
	tx, err := m.db.BeginTx(m.ctx, nil)
	if err != nil {
		m.log.Error("apiIssue", zap.Error(err))
		return config.NewServerError(err)
	}

	var example examplemodel
	err = tx.QueryRowContext(m.ctx, "SELECT id FROM examples WHERE id=?;", id).Scan(&example.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			m.log.Error("notFound", zap.Error(err))
			return config.NewNotfoundError(err)
		}
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	_, err = tx.ExecContext(m.ctx, "DELETE FROM examples WHERE id=?;", example.ID)
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	err = tx.Commit()
	if err != nil {
		m.log.Error("apiIssue", zap.Error(err))
		return config.NewServerError(err)
	}

	return nil
}
