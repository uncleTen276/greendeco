package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Owner       User      `json:"owner" db:"owner_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type CartProduct struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Cart      Cart      `json:"cart" db:"cart_id"`
	Variant   Variant   `json:"variant" db:"variant"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
