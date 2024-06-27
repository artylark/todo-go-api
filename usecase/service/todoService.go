package service

import (
	"github.com/artylark/todo-go-api/domain/model"
	"gorm.io/gorm"
)

type TodoService interface {
	GetAll(todos model.Todos) (model.Todos, error)
	GetById(id int) (model.Todo, error)
}

type todoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) TodoService {
	return &todoService{db: db}
}

func (s *todoService) GetAll(todos model.Todos) (model.Todos, error) {
	err := s.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *todoService) GetById(id int) (model.Todo, error) {
	todo := model.Todo{}
	err := s.db.Where("Id = ?", id).First(&todo).Error
	return todo, err
}
