package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/todo-got/handler"
)

func main() {
	app := echo.New()
	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":8080")
}
