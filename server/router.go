package server

import (
	"go-todo-app/config"

	"github.com/labstack/echo/v4"
)

func NewRouter() (*echo.Echo, error) {
	c := config.GetConfig()
	router := echo.New()

	_ = router.Group("/" + c.GetString("server.version"))

	// homesController := controllers.NewHomesController()
	// version.GET("/", homesController.Index)

	return router, nil
}
