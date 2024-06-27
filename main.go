package main

import (
	"encoding/json"
	"github.com/artylark/todo-go-api/infrastructure"
	"github.com/labstack/echo/v4"
	"time"
)

type Todo struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Todos []Todo

func (Todo) TableName() string {
	return "todo"
}

func main() {
	e := echo.New()

	db := infrastructure.Connect()
	e.GET("/todo/get", func(c echo.Context) error {
		todos := Todos{}
		db.Find(&todos)
		return json.NewEncoder(c.Response()).Encode(todos)
	})
	e.GET("/todo/get/:id", func(c echo.Context) error {
		id := c.Param("id")
		todo := Todo{}
		db.Where("Id = ?", id).First(&todo)
		return json.NewEncoder(c.Response()).Encode(todo)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
