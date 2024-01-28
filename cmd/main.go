package main

import (
	"github.com/Faroukhamadi/usermx/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HTML(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	index := views.Index()
	login := views.Login()

	e.GET("/", func(c echo.Context) error {
		return HTML(c, index)
	})

	e.GET("/login", func(c echo.Context) error {
		return HTML(c, login)
	})
	
	e.Logger.Fatal(e.Start(":42069"))
}
