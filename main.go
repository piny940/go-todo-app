package main

import (
	"flag"
	"go-todo-app/config"
	"go-todo-app/database"
	"go-todo-app/registry"
	"go-todo-app/server"
)

func main() {
	env := flag.String("e", "development", "set environment")
	flag.Parse()

	config.Init(*env)
	database.Init()
	registry.Init(database.GetDB())
	defer database.Close()

	if err := server.Init(); err != nil {
		panic(err)
	}
}
