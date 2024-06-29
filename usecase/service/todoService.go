package service

import (
	"github.com/artylark/todo-go-api/domain/model"
	"github.com/artylark/todo-go-api/usecase/repository"
)

type TodoService interface {
	Create(todo *model.Todo) error
	GetAll() (model.Todos, error)
	GetById(id int) (model.Todo, error)
	Update(todo *model.Todo) error
	Delete(id int) error
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(r repository.TodoRepository) TodoService {
	return &todoService{todoRepository: r}
}

func (s *todoService) Create(todo *model.Todo) error {
	return s.todoRepository.Store(*todo)
}

func (s *todoService) GetAll() (model.Todos, error) {
	return s.todoRepository.FindAll()
}

func (s *todoService) GetById(id int) (model.Todo, error) {
	return s.todoRepository.FindById(id)
}

func (s *todoService) Update(todo *model.Todo) error {
	return s.todoRepository.Update(*todo)
}

func (s *todoService) Delete(id int) error {
	return s.todoRepository.Delete(id)
}
