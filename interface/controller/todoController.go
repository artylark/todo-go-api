package controller

import (
	"github.com/artylark/todo-go-api/domain/model"
	"github.com/artylark/todo-go-api/usecase/service"
)

type TodoController interface {
	GetAllTodos() (model.Todos, error)
	GetTodoById(id int) (model.Todo, error)
}

type todoController struct {
	todoService service.TodoService
}

func NewTodoController(s service.TodoService) TodoController {
	return &todoController{todoService: s}
}

func (c *todoController) GetAllTodos() (model.Todos, error) {
	todos := model.Todos{}
	return c.todoService.GetAll(todos)
}

func (c *todoController) GetTodoById(id int) (model.Todo, error) {
	return c.todoService.GetById(id)
}