package models

import (
	"time"

	"github.com/google/uuid"
)

type Color struct {
	Id        uuid.UUID `db:"id" json:"id"`
	Color     string    `db:"color" json:"color"`
	Name      string    `db:"name" json:"name" `
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateColor struct {
	Color string `db:"color" json:"color" validate:"required,lte=50,iscolor"`
	Name  string `db:"name" json:"name" validate:"required"`
}

type UpdateColor struct {
	ID    uuid.UUID `json:"-" validate:"required,uuid4"`
	Color string    `json:"color" validate:"required,lte=50,iscolor"`
	Name  string    `json:"name" validate:"required"`
}
