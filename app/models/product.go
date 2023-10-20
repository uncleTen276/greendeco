package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Product struct {
	ID          string         `db:"id" json:"id"`
	Category    string         `db:"category_id" json:"category"`
	Name        string         `db:"name" json:"name"`
	IsPublish   bool           `db:"is_publish" json:"is_publish"`
	Size        string         `db:"size" json:"size"`
	Available   bool           `db:"available" json:"available"`
	Type        string         `db:"type" json:"type"`
	Images      pq.StringArray `db:"images" json:"images"`
	Detail      string         `db:"detail" json:"detail"`
	Description *string        `db:"description" json:"description"`
	Light       string         `db:"light" json:"light"`
	Difficulty  string         `db:"difficulty" json:"difficulty"`
	Water       string         `db:"water" json:"water"`
	QrImage     *string        `db:"qr_image" json:"qr_image"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
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
	Water       string    `json:"water" validate:"required,lte=20"`
	Description string    `json:"description"`
}

type UpdateProduct struct {
	ID          string   `json:"-" validate:"uuid4"`
	IsPublish   bool     `json:"is_publish" validate:"required"`
	Available   bool     `json:"available"`
	Size        string   `json:"size" validate:"required,lte=10"`
	Type        string   `json:"type" validate:"required,lte=20"`
	Images      []string `json:"images" validate:"required"`
	Detail      string   `json:"detail" validate:"required"`
	Light       string   `json:"light" validate:"required,lte=50"`
	Difficulty  string   `json:"difficulty" validate:"required"`
	Water       string   `json:"water" validate:"required,lte=20"`
	Description string   `json:"description"`
}

type ProductQuery struct {
	BaseQuery
	Fields ProductQueryField `query:"field"`
}

type ProductQueryField struct {
	Name       string     `query:"name" json:"name"`
	Available  *bool      `query:"available"`
	Category   *uuid.UUID `query:"category_id" json:"category_id" validate:"uuid4"`
	Size       string     `query:"size" json:"size"`
	Type       string     `query:"type" json:"type"`
	Difficulty string     `query:"difficulty" json:"difficulty"`
	Water      string     `query:"water" json:"water"`
}

type ActivedProduct struct {
	ID             string         `db:"id" json:"id"`
	Category       string         `db:"category_id" json:"category"`
	Name           string         `db:"name" json:"name"`
	Price          string         `db:"price" json:"price"`
	Size           string         `db:"size" json:"size"`
	Available      bool           `db:"available" json:"available"`
	Type           string         `db:"type" json:"type"`
	Images         pq.StringArray `db:"images" json:"images"`
	Detail         string         `db:"detail" json:"detail"`
	Description    *string        `db:"description" json:"description"`
	Light          string         `db:"light" json:"light"`
	Difficulty     string         `db:"difficulty" json:"difficulty"`
	Water          string         `db:"water" json:"water"`
	CreatedAt      time.Time      `db:"created_at" json:"created_at"`
	DefaultVariant string         `db:"variant_id" json:"default_variant"`
	Currency       string         `db:"currency" json:"currency"`
}

type UpdateDefaultVariant struct {
	VariantId uuid.UUID `json:"variant" validate:"required,uuid4"`
	ProductId uuid.UUID `json:"-" validate:"required,uuid4"`
}

type DefaultVariant struct {
	VariantId uuid.UUID `json:"variant" db:"variant_id" `
	ProductId uuid.UUID `json:"-" db:"product_id"`
}
