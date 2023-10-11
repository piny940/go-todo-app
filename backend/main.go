package main

import (
	"flag"
	"go-todo-app/config"
	"go-todo-app/db"
	"go-todo-app/registry"
	"go-todo-app/server"
)

func main() {
	env := flag.String("e", "development", "set environment")
	flag.Parse()

	config.Init(*env)
	db.Init()
	registry.Init(db.GetDB())
	defer db.Close()

	if err := server.Init(); err != nil {
		panic(err)
	}
}
