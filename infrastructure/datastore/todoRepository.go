package datastore

import (
	"github.com/artylark/todo-go-api/domain/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Store(todo model.Todo) error
	FindAll() (model.Todos, error)
	FindById(id int) (model.Todo, error)
	Delete(id int) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Store(todo model.Todo) error {
	return r.db.Create(&todo).Error
}

func (r *todoRepository) FindAll() (model.Todos, error) {
	todos := model.Todos{}
	err := r.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) FindById(id int) (model.Todo, error) {
	todo := model.Todo{}
	err := r.db.Where("Id = ?", id).First(&todo).Error
	return todo, err
}

func (r *todoRepository) Delete(id int) error {
	return r.db.Delete(&model.Todo{}, "Id = ?", id).Error
}
