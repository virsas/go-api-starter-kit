package utils

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	vssdb "github.com/virsas/go-mod-db"
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
	}

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

	db, err = vssdb.InitMysqlDB(dbAccountUser, dbAccountPass, dbAccountHost, dbAccountPort, dbAccountName)
	if err != nil {
		return db, err
	}

	return db, nil
}
