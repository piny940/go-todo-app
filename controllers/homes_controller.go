package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomesController struct{}

func NewHomesController() *HomesController {
	return &HomesController{}
}

func (h *HomesController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}
