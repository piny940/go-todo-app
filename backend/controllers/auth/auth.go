package auth

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var sessionsOptions = &sessions.Options{
	Path:     "/",
	MaxAge:   86400 * 7,
	HttpOnly: true,
}

func setSession(c echo.Context, key string, value interface{}) {
	session, _ := session.Get("session", c)
	session.Options = sessionsOptions
	session.Values[key] = value
	session.Save(c.Request(), c.Response())
}
