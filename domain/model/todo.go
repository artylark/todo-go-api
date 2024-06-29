package model

import "time"

type Todo struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Todos []Todo

func (Todo) TableName() string {
	return "todo"
}
