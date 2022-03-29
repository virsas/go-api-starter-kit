package utils

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func InitSQL() (*sql.DB, error) {
	var err error
	var db *sql.DB

	var dbAccountUser string = "root"
	dbAccountUserValue, dbAccountUserPresent := os.LookupEnv("DB_USER")
	if dbAccountUserPresent {
		dbAccountUser = dbAccountUserValue
	}

	var dbAccountPass string = ""
	dbAccountPassValue, dbAccountPassPresent := os.LookupEnv("DB_PASS")
	if dbAccountPassPresent {
		dbAccountPass = dbAccountPassValue
	}

	var dbAccountHost string = "localhost"
	dbAccountHostValue, dbAccountHostPresent := os.LookupEnv("DB_HOST")
	if dbAccountHostPresent {
		dbAccountHost = dbAccountHostValue
	}

	var dbAccountName string = ""
	dbAccountNameValue, dbAccountNamePresent := os.LookupEnv("DB_NAME")
	if dbAccountNamePresent {
		dbAccountName = dbAccountNameValue
	} // initMysqlDB function

	var dbAccountPortStr string = "5432"
	dbAccountPortStrValue, dbAccountPortStrPresent := os.LookupEnv("DB_PORT")
	if dbAccountPortStrPresent {
		dbAccountPortStr = dbAccountPortStrValue
	}
	dbAccountPort, err := strconv.Atoi(dbAccountPortStr)
	if err != nil {
		fmt.Println("Please, configure your environment")
		return db, err
	}

	db, err = initMysqlDB(dbAccountUser, dbAccountPass, dbAccountHost, dbAccountPort, dbAccountName)
	if err != nil {
		return db, err
	}

	return db, nil
}

func initMysqlDB(user string, password string, hostname string, port int, database string) (*sql.DB, error) {
	var err error
	var db *sql.DB

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", user, password, hostname, port, database)
	db, err = sql.Open("mysql", dbSource)
	if err != nil {
		return db, err
	}
	if err = db.Ping(); err != nil {
		return db, err
	}

	openCons, openConsPresent := os.LookupEnv("DB_MAX_OPEN_CONNECTIONS")
	idleCons, idleConsPresent := os.LookupEnv("DB_MAX_IDLE_CONNECTIONS")

	db.SetMaxOpenConns(25)
	if openConsPresent {
		openConsInt, err := strconv.Atoi(openCons)
		if err != nil {
			return db, err
		}
		db.SetMaxOpenConns(openConsInt)
	}

	db.SetMaxIdleConns(25)
	if idleConsPresent {
		idleConsInt, err := strconv.Atoi(idleCons)
		if err != nil {
			return db, err
		}
		db.SetMaxIdleConns(idleConsInt)
	}

	return db, err
}
