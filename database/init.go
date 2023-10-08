package database

import (
	"database/sql"
	"go-todo-app/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	c := config.GetConfig()
	var err error
	db, err = sql.Open(c.GetString("db.provider"), c.GetString("db.url"))
	if err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func Close() {
	db.Close()
}
