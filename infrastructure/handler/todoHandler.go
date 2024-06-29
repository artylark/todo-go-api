package handler

import (
	"encoding/json"
	"github.com/artylark/todo-go-api/domain/model"
	"github.com/artylark/todo-go-api/interface/controller"
	"github.com/labstack/echo/v4"
	"strconv"
)

type TodoHandler interface {
	CreateTodo(c echo.Context) error
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoHandler struct {
	todoController controller.TodoController
}

func NewTodoHandler(c controller.TodoController) TodoHandler {
	return &todoHandler{todoController: c}
}

func (h *todoHandler) CreateTodo(c echo.Context) error {
	todo := &model.Todo{}
	if err := json.NewDecoder(c.Request().Body).Decode(&todo); err != nil {
		return err
	}
	return h.todoController.CreateTodo(todo)
}

func (h *todoHandler) GetAllTodos(c echo.Context) error {
	todos, err := h.todoController.GetAllTodos()
	if err != nil {
		return err
	}
	return json.NewEncoder(c.Response()).Encode(todos)
}

func (h *todoHandler) GetTodoById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	todo, err := h.todoController.GetTodoById(id)
	if err != nil {
		return err
	}
	return json.NewEncoder(c.Response()).Encode(todo)
}

func (h *todoHandler) UpdateTodo(c echo.Context) error {
	todo := &model.Todo{}
	if err := json.NewDecoder(c.Request().Body).Decode(&todo); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	todo.Id = id
	return h.todoController.UpdateTodo(todo)
}

func (h *todoHandler) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	return h.todoController.DeleteTodo(id)
}
