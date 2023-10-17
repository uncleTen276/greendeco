package models

import (
	"time"

	"github.com/google/uuid"
)

type RecommendProduct struct {
	ProductId   uuid.UUID `json:"product" validate:"required,uuid4" db:"product_id"`
	RecommendId uuid.UUID `json:"recommend" validate:"required,uuid4" db:"recommend_product"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type CreateRecommend struct {
	ProductId   uuid.UUID `json:"product" validate:"required,uuid4" db:"product_id"`
	RecommendId uuid.UUID `json:"recommend" validate:"required,uuid4" db:"recommend_product"`
}
