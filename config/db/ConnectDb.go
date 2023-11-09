package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

// connect database
func ConnectDb() {
	var err error
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "HoangLuat123"
	dbName := "thuchanh"

	DB, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		log.Fatal(err)
		panic("Failed to create a connection to database")
	}

}

// close database
func CloseDb() {
	ConnectDb()
	if err != nil {
		panic("Failed to close a connection to database")
	}
	DB.Close()
}