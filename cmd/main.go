package main

import (
	"github.com/blakeb21/go_playground/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")
}
