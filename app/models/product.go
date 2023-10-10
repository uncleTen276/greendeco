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
	CategoryId  uuid.UUID `json:"category_id"`
	Name        string    `json:"name"`
	Size        string    `json:"size"`
	Type        string    `json:"type"`
	Images      []string  `json:"images"`
	Detail      string    `json:"detail"`
	Light       string    `json:"light"`
	Difficulty  string    `json:"difficulty"`
	Warter      string    `json:"warter"`
	QrImage     string    `json:"qr_image"`
	Description string    `json:"description"`
}
