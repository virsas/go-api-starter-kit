package db

import (
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func PostgresMigrate(db *sql.DB, path string) error {
	var err error

	var migrationTable string = "golang_migrations"
	migrationTableValue, migrationTablePresent := os.LookupEnv("DB_MIGRATION_TABLE")
	if migrationTablePresent {
		migrationTable = migrationTableValue
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: migrationTable,
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return err
	}

	return nil
}

func MysqlMigrate(db *sql.DB, path string) error {
	var err error

	var migrationTable string = "golang_migrations"
	migrationTableValue, migrationTablePresent := os.LookupEnv("DB_MIGRATION_TABLE")
	if migrationTablePresent {
		migrationTable = migrationTableValue
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: migrationTable,
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(path, "mysql", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return err
	}

	return nil
}
