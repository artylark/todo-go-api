package router

import (
	"encoding/json"
	"github.com/artylark/todo-go-api/domain/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(e *echo.Echo, db *gorm.DB) {
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
}
