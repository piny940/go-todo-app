package db

import (
	"database/sql"
	"go-todo-app/config"

	_ "github.com/lib/pq"
)

type DB struct {
	Client *sql.DB
}

var db *DB

func Init() {
	c := config.GetConfig()
	client, err := sql.Open(c.GetString("db.provider"), c.GetString("db.dsn"))
	if err != nil {
		panic(err)
	}
	db = &DB{Client: client}
}

func GetDB() *DB {
	return db
}

func Close() {
	db.Client.Close()
}
