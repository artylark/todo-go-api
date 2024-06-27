package handler

import (
	"encoding/json"
	"github.com/artylark/todo-go-api/domain/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TodoHandler interface {
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
}

type todoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) TodoHandler {
	return &todoHandler{db: db}
}

func (h *todoHandler) GetAllTodos(c echo.Context) error {
	todos := model.Todos{}
	h.db.Find(&todos)
	return json.NewEncoder(c.Response()).Encode(todos)
}

func (h *todoHandler) GetTodoById(c echo.Context) error {
	id := c.Param("id")
	todo := model.Todo{}
	h.db.Where("Id = ?", id).First(&todo)
	return json.NewEncoder(c.Response()).Encode(todo)
}
