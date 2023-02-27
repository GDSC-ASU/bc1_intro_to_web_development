package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var connectionInstance *sql.DB = nil

func GetInstance() *sql.DB {
	if connectionInstance == nil {
		var err error
		// DB_USERNAME is root by default
		// DB_PASSWORD is what you've set when installing the database
		// DB_NAME is the database you've created using `CREATE DATABASE whateverName`
		// connectionInstance, err = sql.Open("mysql", "DB_USERNAME:DB_PASSWORD@tcp(localhost:3306)/DB_NAME?parseTime=True")
		connectionInstance, err = sql.Open("mysql", "root:1234@tcp(localhost:3306)/ttt?parseTime=True")
		if err != nil {
			panic(err)
		}
	}
	return connectionInstance
}
