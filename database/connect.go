package database

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func DBConn() (*sql.DB, error) {
	dbDriver 	:= "mysql"
	dbUser		:= "root"
	dbPass		:= "HoangLuat123"
	dbName 		:= "thuchanh"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}