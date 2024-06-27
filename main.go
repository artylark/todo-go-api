package main

import (
	"github.com/artylark/todo-go-api/infrastructure/datastore"
	"github.com/artylark/todo-go-api/infrastructure/handler"
	"github.com/artylark/todo-go-api/infrastructure/router"
	"github.com/artylark/todo-go-api/interface/controller"
	"github.com/artylark/todo-go-api/usecase/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := datastore.Connect()
	r := datastore.NewTodoRepository(db)
	s := service.NewTodoService(r)
	c := controller.NewTodoController(s)
	h := handler.NewTodoHandler(c)
	router.Router(e, h)
	e.Logger.Fatal(e.Start(":8080"))
}
