package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Owner       uuid.UUID `json:"owner" db:"owner_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type CartProduct struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Cart      Cart      `json:"cart" db:"cart_id"`
	Variant   Variant   `json:"variant" db:"variant"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateCart struct {
	Owner       uuid.UUID `json:"-" validate:"required,uuid4"`
	Description string    `json:"description"`
}

type CreateCartProduct struct {
	Cart     uuid.UUID `json:"cart_id" validate:"required,uuid4"`
	Variant  uuid.UUID `json:"variant_id" validate:"required,uuid4"`
	Quantity int       `json:"quantity" validate:"required,numeric,gte=1"`
}
