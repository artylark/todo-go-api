package repository

import "github.com/artylark/todo-go-api/domain/model"

type TodoRepository interface {
	FindAll() (model.Todos, error)
	FindById(id int) (model.Todo, error)
}
