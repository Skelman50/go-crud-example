package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "admin:130588@tcp(127.0.0.1:3306)/myTestDB")
	if err != nil {
		fmt.Println(err)
	}
	Db.SetMaxIdleConns(75)
	Db.SetMaxOpenConns(99)
}
