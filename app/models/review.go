package models

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID `db:"id" json:"id"`
	UserId    uuid.UUID `db:"user_id" json:"user_id"`
	ProductId uuid.UUID `db:"product_id" json:"product_id"`
	Content   string    `db:"content" json:"content"`
	Star      int       `db:"star" json:"star"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateReview struct {
	UserId    uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid4"`
	ProductId uuid.UUID `db:"product_id" json:"product_id" validate:"required,uuid4"`
	Content   string    `db:"content" json:"content" validate:"omitempty"`
	Star      int       `db:"star" json:"star" validate:"omitempty"`
}

type ReviewQuery struct {
	BaseQuery
	Star   int        `db:"star" json:"star" query:"star"`
	UserId *uuid.UUID `db:"user_id" json:"user_id" query:"user_id"`
}
