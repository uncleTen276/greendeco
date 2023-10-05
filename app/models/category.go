package models

import "time"

type Category struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewCategory() *Category {
	return &Category{}
}

type CreateCategory struct {
	Name string `json:"name" validate:"required,lte=100"`
}

type UpdateCategory struct {
	ID   string `json:"-"`
	Name string `json:"name" validate:"required,lte=100"`
}
