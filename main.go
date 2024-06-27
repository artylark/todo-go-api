package main

import (
	"encoding/json"
	"github.com/artylark/todo-go-api/domain/model"
	"github.com/artylark/todo-go-api/infrastructure"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := infrastructure.Connect()
	e.GET("/todo/get", func(c echo.Context) error {
		todos := model.Todos{}
		db.Find(&todos)
		return json.NewEncoder(c.Response()).Encode(todos)
	})
	e.GET("/todo/get/:id", func(c echo.Context) error {
		id := c.Param("id")
		todo := model.Todo{}
		db.Where("Id = ?", id).First(&todo)
		return json.NewEncoder(c.Response()).Encode(todo)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
