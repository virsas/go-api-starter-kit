package example

import (
	"context"
	"database/sql"
	"fmt"
	"go-api-starter-kit/utils/logger"
	"go-api-starter-kit/utils/vars"
	"time"
)

type model struct {
	db  *sql.DB
	log logger.LoggerHandler
}

func newModel(db *sql.DB, log logger.LoggerHandler) *model {
	return &model{db: db, log: log}
}

// ExampleInput struct
type ExampleInput struct {
	Name *string `json:"name" validate:"required,alphanum,min=2,max=255"`
}

// Example struct
type Example struct {
	ID *int64 `json:"id"`
	ExampleInput
	Updatedat *time.Time `json:"updatedat"`
	Createdat *time.Time `json:"createdat"`
}

const exampleGetQuery string = "SELECT id FROM examples WHERE id=$1 and account_id=$2;"
const exampleListQuery string = "SELECT id, name, updated_at, created_at FROM examples WHERE account_id=$1 ORDER BY id DESC;"
const exampleShowQuery string = "SELECT id, name, updated_at, created_at FROM examples WHERE id=$1 and account_id=$2;"
const exampleInsertQuery string = "INSERT INTO examples (name, account_id, user_id, updated_at, created_at) VALUES ($1,$2,$3,$4,$5);"
const exampleUpdateQuery string = "UPDATE examples SET name=$1, updated_at=$2 WHERE id=$3;"
const exampleDeleteQuery string = "DELETE FROM examples WHERE id=$1;"

func (m *model) list(aid int64, uid int64) ([]Example, error) {
	var example Example
	var examples []Example = []Example{}

	exampleListQuery, err := m.db.Prepare(exampleListQuery)
	if err != nil {
		m.log.Error(err.Error())
		return examples, vars.StatusDBError(err)
	}

	exampleListQueryExec, err := exampleListQuery.Query(aid)
	if err != nil {
		m.log.Error(err.Error())
		return examples, vars.StatusDBError(err)
	}
	defer exampleListQueryExec.Close()

	for exampleListQueryExec.Next() {
		err = exampleListQueryExec.Scan(&example.ID, &example.Name, &example.Updatedat, &example.Createdat)
		if err != nil {
			return examples, vars.StatusDBError(err)
		}

		examples = append(examples, example)
	}

	return examples, nil
}

func (m *model) create(c context.Context, model ExampleInput, aid int64, uid int64) error {
	tx, err := m.db.BeginTx(c, nil)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusServerError(err)
	}

	exampleInsertQuery, err := m.db.PrepareContext(c, exampleInsertQuery)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	fmt.Println(*model.Name)
	_, err = exampleInsertQuery.ExecContext(c, model.Name, aid, uid, time.Now(), time.Now())
	if err != nil {
		tx.Rollback()
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	err = tx.Commit()
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusServerError(err)
	}

	return nil
}

func (m *model) show(id int64, aid int64, uid int64) (Example, error) {
	var example Example

	exampleShowQuery, err := m.db.Prepare(exampleShowQuery)
	if err != nil {
		m.log.Error(err.Error())
		return example, vars.StatusDBError(err)
	}

	err = exampleShowQuery.QueryRow(id, aid).Scan(&example.ID, &example.Name, &example.Updatedat, &example.Createdat)
	if err != nil {
		if err == sql.ErrNoRows {
			m.log.Error(err.Error())
			return example, vars.StatusNotfoundError(err)
		}
		m.log.Error(err.Error())
		return example, vars.StatusDBError(err)
	}

	return example, nil
}

func (m *model) update(c context.Context, id int64, model ExampleInput, aid int64, uid int64) error {
	tx, err := m.db.BeginTx(c, nil)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusServerError(err)
	}

	exampleGetQuery, err := tx.PrepareContext(c, exampleGetQuery)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	var example Example
	err = exampleGetQuery.QueryRowContext(c, id, aid).Scan(&example.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			m.log.Error(err.Error())
			return vars.StatusNotfoundError(err)
		}
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	exampleUpdateQuery, err := tx.PrepareContext(c, exampleUpdateQuery)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	_, err = exampleUpdateQuery.ExecContext(c, model.Name, time.Now(), example.ID)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	err = tx.Commit()
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusServerError(err)
	}

	return nil
}

func (m *model) delete(c context.Context, id int64, aid int64, uid int64) error {
	tx, err := m.db.BeginTx(c, nil)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusServerError(err)
	}

	exampleGetQuery, err := tx.PrepareContext(c, exampleGetQuery)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	var example Example
	err = exampleGetQuery.QueryRowContext(c, id, aid).Scan(&example.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			m.log.Error(err.Error())
			return vars.StatusNotfoundError(err)
		}
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	exampleDeleteQuery, err := tx.PrepareContext(c, exampleDeleteQuery)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	_, err = exampleDeleteQuery.ExecContext(c, example.ID)
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusDBError(err)
	}

	err = tx.Commit()
	if err != nil {
		m.log.Error(err.Error())
		return vars.StatusServerError(err)
	}

	return nil
}
