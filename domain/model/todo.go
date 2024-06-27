package model

import "time"

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
