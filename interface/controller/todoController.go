package controller

import (
	"github.com/artylark/todo-go-api/domain/model"
	"gorm.io/gorm"
)

type TodoController interface {
	GetAllTodos() (model.Todos, error)
	GetTodoById(id int) (model.Todo, error)
}

type todoController struct {
	db *gorm.DB
}

func NewTodoController(db *gorm.DB) TodoController {
	return &todoController{db: db}
}

func (c *todoController) GetAllTodos() (model.Todos, error) {
	todos := model.Todos{}
	err := c.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (c *todoController) GetTodoById(id int) (model.Todo, error) {
	todo := model.Todo{}
	err := c.db.Where("Id = ?", id).First(&todo).Error
	return todo, err
}
