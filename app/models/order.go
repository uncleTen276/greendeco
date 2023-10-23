package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID              uuid.UUID `json:"id" db:"id"`
	OwnerId         uuid.UUID `json:"owner_id" db:"owner_id"`
	PaidAt          time.Time `json:"paid_at" db:"paid_at"`
	UserName        string    `json:"user_name" db:"user_name"`
	UserEmail       string    `json:"user_email" db:"user_email"`
	ShippingAddress string    `json:"shipping_address" db:"shipping_address"`
	UserPhoneNumber string    `json:"user_phone_number" db:"user_phoneNumber"`
	State           string    `json:"state" db:"state"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type OrderProduct struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CartID       uuid.UUID `json:"cart_id" db:"cart_id"`
	VariantId    uuid.UUID `json:"variant_id" db:"variant_id"`
	VariantName  string    `json:"variant_name" db:"variant_name"`
	VariantPrice string    `json:"variant_price" db:"variant_price"`
	ActualPrice  string    `json:"actual_price" db:"actual_price"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type CreateOrder struct {
	Owner uuid.UUID `json:"owner_id"`
}

type CreateOrderProduct struct{}
