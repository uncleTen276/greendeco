package models

import (
	"time"

	"github.com/google/uuid"
)

type Variant struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Available   bool      `json:"available" db:"available"`
	Product     uuid.UUID `json:"product" db:"product"`
	Name        string    `json:"name" db:"name"`
	Color       string    `json:"color" db:"color"`
	ColorName   string    `json:"color_name" db:"color_name"`
	Price       string    `json:"price" db:"price"`
	Image       string    `json:"image" db:"image"`
	Description string    `json:"description" db:"description"`
	Currency    string    `json:"currency" db:"currency"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type CreateVariant struct {
	IsDefault   bool      `json:"is_default"`
	Available   bool      `json:"available" db:"available"`
	ProductId   uuid.UUID `json:"product_id" validate:"required"`
	Name        string    `json:"name" validate:"required,gte=3,lte=50"`
	Color       string    `json:"color" validate:"required,gte=3,lte=50,iscolor"`
	ColorName   string    `json:"color_name" validate:"required"`
	Price       int       `json:"price" validate:"required"`
	Currency    string    `json:"currency" validate:"required,iso4217"`
	Image       string    `json:"image" validate:"required,url"`
	Description string    `json:"description"`
}

type UpdateVariant struct {
	ID          uuid.UUID `json:"-,omitempty" db:"id"`
	IsDefault   bool      `json:"is_default"`
	Available   bool      `json:"available" db:"available"`
	ProductId   uuid.UUID `json:"product_id" validate:"required"`
	Name        string    `json:"name" validate:"required,gte=3,lte=50"`
	Color       string    `json:"color" validate:"required,gte=3,lte=50,iscolor"`
	ColorName   string    `json:"color_name" validate:"required"`
	Price       int       `json:"price" validate:"required"`
	Currency    string    `json:"currency" validate:"required,iso4217"`
	Image       string    `json:"image" validate:"required,url"`
	Description string    `json:"description"`
}
