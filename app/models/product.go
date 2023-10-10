package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          string `db:"id"`
	Category    `db:"category"`
	Name        string    `db:"name"`
	IsPublish   bool      `db:"is_publish"`
	Size        string    `db:"size"`
	Type        string    `db:"type"`
	Images      []string  `db:"images"`
	Detail      string    `db:"detail"`
	Description string    `db:"description"`
	Light       string    `db:"light"`
	Difficulty  string    `db:"difficulty"`
	Warter      string    `db:"warter"`
	QrImage     string    `db:"qr_image"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewProduct() *Product {
	return &Product{}
}

type CreateProduct struct {
	CategoryId  uuid.UUID `json:"category_id" validate:"required,uuid4"`
	Name        string    `json:"name" validate:"required"`
	Size        string    `json:"size" validate:"required,lte=10"`
	Type        string    `json:"type" validate:"required,lte=20"`
	Images      []string  `json:"images" validate:"required"`
	Detail      string    `json:"detail" validate:"required"`
	Light       string    `json:"light" validate:"required,lte=50"`
	Difficulty  string    `json:"difficulty" validate:"required"`
	Warter      string    `json:"warter" validate:"required,lte=20"`
	Description string    `json:"description"`
}

type UpdateProduct struct {
	ID          string   `json:"-" validate:"uuid4"`
	IsPublish   bool     `json:"is_publish" validate:"required"`
	Size        string   `json:"size" validate:"required,lte=10"`
	Type        string   `json:"type" validate:"required,lte=20"`
	Images      []string `json:"images" validate:"required"`
	Detail      string   `json:"detail" validate:"required"`
	Light       string   `json:"light" validate:"required,lte=50"`
	Difficulty  string   `json:"difficulty" validate:"required"`
	Warter      string   `json:"warter" validate:"required,lte=20"`
	Description string   `json:"description"`
}
