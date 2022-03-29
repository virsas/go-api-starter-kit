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
}

func newModel(db *sql.DB, log *zap.Logger) *model {
	return &model{db: db, log: log}
}

// ExampleInput struct
type ExampleInput struct {
	Name *string `json:"name" validate:"required,min=2,max=255"`
}

// Example struct
type Example struct {
	ID *int64 `json:"id"`
	ExampleInput
	Updatedat *time.Time `json:"updatedat"`
	Createdat *time.Time `json:"createdat"`
}

func (m *model) list(aid int) ([]Example, error) {
	var example Example
	var examples []Example = []Example{}

	exampleListQuery, err := m.db.Prepare("SELECT id, name, updated_at, created_at FROM examples ORDER BY id DESC;")
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return examples, config.NewDBError(err)
	}

	exampleListQueryExec, err := exampleListQuery.Query()
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return examples, config.NewDBError(err)
	}
	defer exampleListQueryExec.Close()

	for exampleListQueryExec.Next() {
		err = exampleListQueryExec.Scan(&example.ID, &example.Name, &example.Updatedat, &example.Createdat)
		if err != nil {
			return examples, config.NewDBError(err)
		}

		examples = append(examples, example)
	}

	return examples, nil
}

func (m *model) create(c context.Context, model ExampleInput, aid int) error {
	tx, err := m.db.BeginTx(c, nil)
	if err != nil {
		m.log.Error("apiIssue", zap.Error(err))
		return config.NewServerError(err)
	}

	exampleInsertQuery, err := m.db.PrepareContext(c, "INSERT INTO examples (name, updated_at, created_at) VALUES (?,?,?);")
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	_, err = exampleInsertQuery.ExecContext(c, model.Name, time.Now(), time.Now())
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

func (m *model) show(id int64, aid int) (Example, error) {
	var example Example

	exampleShowQuery, err := m.db.Prepare("SELECT id, name, updated_at, created_at FROM examples WHERE id=?;")
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return example, config.NewDBError(err)
	}

	err = exampleShowQuery.QueryRow(id).Scan(&example.ID, &example.Name, &example.Updatedat, &example.Createdat)
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

func (m *model) update(c context.Context, id int64, model ExampleInput, aid int) error {
	tx, err := m.db.BeginTx(c, nil)
	if err != nil {
		m.log.Error("apiIssue", zap.Error(err))
		return config.NewServerError(err)
	}

	exampleGetQuery, err := tx.PrepareContext(c, "SELECT id FROM examples WHERE id=?;")
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	var example Example
	err = exampleGetQuery.QueryRowContext(c, id).Scan(&example.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			m.log.Error("notFound", zap.Error(err))
			return config.NewNotfoundError(err)
		}
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	exampleUpdateQuery, err := tx.PrepareContext(c, "UPDATE examples SET name=?, updated_at=? WHERE id=?;")
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	_, err = exampleUpdateQuery.ExecContext(c, model.Name, time.Now(), example.ID)
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

func (m *model) delete(c context.Context, id int64, aid int) error {
	tx, err := m.db.BeginTx(c, nil)
	if err != nil {
		m.log.Error("apiIssue", zap.Error(err))
		return config.NewServerError(err)
	}

	exampleGetQuery, err := tx.PrepareContext(c, "SELECT id FROM examples WHERE id=?;")
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	var example Example
	err = exampleGetQuery.QueryRowContext(c, id).Scan(&example.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			m.log.Error("notFound", zap.Error(err))
			return config.NewNotfoundError(err)
		}
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	exampleDeleteQuery, err := tx.PrepareContext(c, "DELETE FROM examples WHERE id=?;")
	if err != nil {
		m.log.Error("dbError", zap.Error(err))
		return config.NewDBError(err)
	}

	_, err = exampleDeleteQuery.ExecContext(c, example.ID)
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
