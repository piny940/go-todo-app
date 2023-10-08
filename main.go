package main

import (
	"flag"
	"go-todo-app/config"
	"go-todo-app/database"
)

func main() {
	env := flag.String("e", "development", "set environment")
	flag.Parse()

	config.Init(*env)
	database.Init()
	defer database.Close()
}
