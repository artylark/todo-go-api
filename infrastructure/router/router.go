package router

import (
	"github.com/artylark/todo-go-api/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, h handler.TodoHandler) {
	e.GET("/todo/get", h.GetAllTodos)
	e.GET("/todo/get/:id", h.GetTodoById)
}
