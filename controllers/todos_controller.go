package controllers

import (
	"go-todo-app/registry"
	"net/http"

	"github.com/labstack/echo/v4"
)

type todosController struct{}

func NewTodosController() *todosController {
	return &todosController{}
}

func (t *todosController) Index(c echo.Context) error {
	registry := registry.GetRegistry()
	todos, err := registry.TodoUseCase().List()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"todos": todos,
	})
}
