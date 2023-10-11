package models

import (
	"time"

	"github.com/google/uuid"
)

type Variant struct {
	ID          uuid.UUID `json:"id" db:"id"`
	IsDefault   bool      `json:"is_default" db:"is_default"`
	Product     `json:"product" db:"product"`
	Name        string    `json:"name" db:"name"`
	Color       string    `json:"color" db:"color"`
	Price       string    `json:"price" db:"price"`
	Image       string    `json:"image" db:"image"`
	Description string    `json:"description" db:"description"`
	Currency    string    `json:"currency" db:"currency"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type CreateVariant struct {
	IsDefault   bool      `json:"is_default" validate:"required"`
	ProductId   uuid.UUID `json:"product_id" validate:"required"`
	Name        string    `json:"name" validate:"required,gte=3,lte=50"`
	Color       string    `json:"color" validate:"required,gte=3,lte=50"`
	Price       int       `json:"price" validate:"required"`
	Currency    string    `json:"currency" validate:"required,iso4217"`
	Image       string    `json:"image" validate:"required,url"`
	Description string    `json:"description"`
}
