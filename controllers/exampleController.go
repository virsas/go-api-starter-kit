package controllers

import (
	"context"
	"database/sql"
	"go-api-starter-kit/models"

	"go.uber.org/zap"
)

func ExampleList(db *sql.DB, log *zap.Logger, ctx context.Context) (int, string, map[string]interface{}) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue", nil
	}

	var example models.Example
	var examples []models.Example = []models.Example{}
	examplesQuery, err := tx.QueryContext(ctx, "SELECT id, name, updated_at, created_at FROM examples ORDER BY id DESC;")
	if err != nil {
		log.Error("dbError", zap.Error(err))
		return 400, "dbError", nil
	}
	defer examplesQuery.Close()

	for examplesQuery.Next() {
		err = examplesQuery.Scan(&example.ID, &example.Name, &example.Updatedat, &example.Createdat)
		if err != nil {
			log.Error("dbError", zap.Error(err))
			return 400, "dbError", nil
		}

		examples = append(examples, example)
	}

	err = tx.Commit()
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue", nil
	}

	payload := map[string]interface{}{
		"examples": examples,
	}

	return 200, "OK", payload
}

func ExampleCreate(model models.ExampleReq, db *sql.DB, log *zap.Logger, ctx context.Context) (int, string) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue"
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO examples (name) VALUES (?);", model.Name)
	if err != nil {
		tx.Rollback()
		log.Error("dbError", zap.Error(err))
		return 400, "dbError"
	}

	err = tx.Commit()
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue"
	}

	return 200, "OK"
}

func ExampleShow(id int64, db *sql.DB, log *zap.Logger, ctx context.Context) (int, string, map[string]interface{}) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue", nil
	}

	var example models.Example
	err = tx.QueryRowContext(ctx, "SELECT id, name, updated_at, created_at FROM examples WHERE id=?;", id).Scan(&example.ID, &example.Name, &example.Updatedat, &example.Createdat)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("notFound", zap.Error(err))
			return 404, "notFound", nil
		}
		log.Error("dbError", zap.Error(err))
		return 400, "dbError", nil
	}

	err = tx.Commit()
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue", nil
	}

	payload := map[string]interface{}{
		"example": example,
	}

	return 200, "OK", payload
}

func ExampleUpdate(id int64, model models.ExampleReq, db *sql.DB, log *zap.Logger, ctx context.Context) (int, string) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue"
	}

	var example models.Example
	err = tx.QueryRowContext(ctx, "SELECT id FROM examples WHERE id=?;", id).Scan(&example.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("notFound", zap.Error(err))
			return 404, "notFound"
		}
		log.Error("dbError", zap.Error(err))
		return 400, "dbError"
	}

	_, err = tx.ExecContext(ctx, "UPDATE examples SET name=? WHERE id=?;", model.Name, example.ID)
	if err != nil {
		log.Error("dbError", zap.Error(err))
		return 400, "dbError"
	}

	err = tx.Commit()
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue"
	}

	return 200, "OK"
}

func ExampleDelete(id int64, db *sql.DB, log *zap.Logger, ctx context.Context) (int, string) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue"
	}

	var example models.Example
	err = tx.QueryRowContext(ctx, "SELECT id FROM examples WHERE id=?;", id).Scan(&example.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("notFound", zap.Error(err))
			return 404, "notFound"
		}
		log.Error("dbError", zap.Error(err))
		return 400, "dbError"
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM examples WHERE id=?;", example.ID)
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue"
	}

	err = tx.Commit()
	if err != nil {
		log.Error("apiIssue", zap.Error(err))
		return 400, "apiIssue"
	}

	return 200, "OK"
}
