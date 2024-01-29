package main

import (
	"fmt"

	"github.com/Faroukhamadi/usermx/utils"
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
	login := views.Login(utils.NewPage())
	signup := views.Signup()

	page := utils.NewPage()

	e.GET("/", func(c echo.Context) error {
		return HTML(c, index)
	})

	e.GET("/login", func(c echo.Context) error {
		return HTML(c, login)
	})

	e.POST("/login", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		fmt.Println("email:", email)
		fmt.Println("password:", password)

		if !page.Data.UserExists(email) {
			loginForm := views.LoginForm(page)
			page.Form.Errors["email"] = "User does not exist"
			c.Response().Writer.WriteHeader(422)
			return HTML(c, loginForm)
		}
		return HTML(c, login)
	})

	e.GET("/signup", func(c echo.Context) error {
		return HTML(c, signup)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
