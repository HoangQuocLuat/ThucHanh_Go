package database

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func DBConn() (*sql.DB, error) {
	dbDriver 	:= "mysql"
	dbUser		:= "root"
	dbPass		:= "123"
	dbName 		:= "thuchanh"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(ThucHanh:3306)/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}