package router

import (
	"github.com/artylark/todo-go-api/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, h handler.TodoHandler) {
	e.POST("/todo", h.CreateTodo)
	e.GET("/todo", h.GetAllTodos)
	e.GET("/todo/:id", h.GetTodoById)
	e.DELETE("/todo/:id", h.DeleteTodo)
}
