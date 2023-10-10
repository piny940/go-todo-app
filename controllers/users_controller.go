package controllers

import (
	"go-todo-app/domain"
	"go-todo-app/registry"
	"net/http"

	"github.com/labstack/echo/v4"
)

type usersController struct{}

func NewUsersController() *usersController {
	return &usersController{}
}

func (u *usersController) Index(c echo.Context) error {
	registry := registry.GetRegistry()
	users, err := registry.UserUseCase().List()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"users": users,
	})
}

func (u *usersController) Create(c echo.Context) error {
	registry := registry.GetRegistry()
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := registry.UserUseCase().SignUp(
		domain.UserEmail(email), domain.UserRawPassword(password),
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}
