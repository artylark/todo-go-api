package controller

import (
	"github.com/artylark/todo-go-api/domain/model"
	"github.com/artylark/todo-go-api/usecase/service"
)

type TodoController interface {
	CreateTodo(todo *model.Todo) error
	GetAllTodos() (model.Todos, error)
	GetTodoById(id int) (model.Todo, error)
}

type todoController struct {
	todoService service.TodoService
}

func NewTodoController(s service.TodoService) TodoController {
	return &todoController{todoService: s}
}

func (c *todoController) CreateTodo(todo *model.Todo) error {
	return c.todoService.Create(todo)
}

func (c *todoController) GetAllTodos() (model.Todos, error) {
	return c.todoService.GetAll()
}

func (c *todoController) GetTodoById(id int) (model.Todo, error) {
	return c.todoService.GetById(id)
}
