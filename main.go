package main

import (
	"github.com/artylark/todo-go-api/infrastructure/datastore"
	"github.com/artylark/todo-go-api/infrastructure/handler"
	"github.com/artylark/todo-go-api/infrastructure/router"
	"github.com/artylark/todo-go-api/interface/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := datastore.Connect()

	c := controller.NewTodoController(db)

	h := handler.NewTodoHandler(c)

	router.Router(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
