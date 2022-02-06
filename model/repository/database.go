package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"todo-app", "todo-password", "go_rest_mvc_test_db:3306", "todo",
	)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}
