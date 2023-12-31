package db

import (
	"database/sql"
	"fmt"
	"go-todo-app/config"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	Client *sql.DB
}

var db *DB

func Init() {
	c := config.GetConfig()
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
	client, err := sql.Open(c.GetString("db.provider"), dsn)
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
