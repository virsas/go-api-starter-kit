package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	var err error
	var db *sql.DB

	var dbUser string = "root"
	dbUserValue, dbUserPresent := os.LookupEnv("DB_USER")
	if dbUserPresent {
		dbUser = dbUserValue
	}

	var dbPass string = ""
	dbPassValue, dbPassPresent := os.LookupEnv("DB_PASS")
	if dbPassPresent {
		dbPass = dbPassValue
	}

	var dbHost string = "localhost"
	dbHostValue, dbHostPresent := os.LookupEnv("DB_HOST")
	if dbHostPresent {
		dbHost = dbHostValue
	}

	var dbName string = "test"
	dbNameValue, dbNamePresent := os.LookupEnv("DB_NAME")
	if dbNamePresent {
		dbName = dbNameValue
	}

	var dbPort string = "3306"
	dbPortValue, dbPortPresent := os.LookupEnv("DB_PORT")
	if dbPortPresent {
		dbPort = dbPortValue
	}

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dbSource)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	openCons, openConsPresent := os.LookupEnv("DB_MAX_OPEN_CONNECTIONS")
	idleCons, idleConsPresent := os.LookupEnv("DB_MAX_IDLE_CONNECTIONS")

	db.SetMaxOpenConns(25)
	if openConsPresent {
		openConsInt, err := strconv.Atoi(openCons)
		if err != nil {
			return nil, err
		}
		db.SetMaxOpenConns(openConsInt)
	}

	db.SetMaxIdleConns(25)
	if idleConsPresent {
		idleConsInt, err := strconv.Atoi(idleCons)
		if err != nil {
			return nil, err
		}
		db.SetMaxIdleConns(idleConsInt)
	}

	return db, nil
}
