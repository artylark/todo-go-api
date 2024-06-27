package service

import (
	"github.com/artylark/todo-go-api/domain/model"
	"github.com/artylark/todo-go-api/usecase/repository"
)

type TodoService interface {
	GetAll() (model.Todos, error)
	GetById(id int) (model.Todo, error)
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(r repository.TodoRepository) TodoService {
	return &todoService{todoRepository: r}
}

func (s *todoService) GetAll() (model.Todos, error) {
	return s.todoRepository.FindAll()
}

func (s *todoService) GetById(id int) (model.Todo, error) {
	return s.todoRepository.FindById(id)
}
