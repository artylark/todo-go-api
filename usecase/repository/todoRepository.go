package repository

import "github.com/artylark/todo-go-api/domain/model"

type TodoRepository interface {
	Store(todo model.Todo) error
	FindAll() (model.Todos, error)
	FindById(id int) (model.Todo, error)
	Update(todo model.Todo) error
	Delete(id int) error
}
