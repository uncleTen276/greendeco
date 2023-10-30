package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusDraft      = "draft"
	StatusProcessing = "processing"
	StatusCompleted  = "completed"
	StatusCancelled  = "cancelled"
)

var StatusTable = map[string][]string{
	StatusDraft:      {StatusProcessing, StatusCancelled},
	StatusProcessing: {StatusCompleted, StatusCancelled},
	StatusCompleted:  {},
	StatusCancelled:  {},
}

type Order struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	OwnerId         uuid.UUID  `json:"owner_id" db:"owner_id"`
	UserName        string     `json:"user_name" db:"user_name"`
	UserEmail       string     `json:"user_email" db:"user_email"`
	ShippingAddress string     `json:"shipping_address" db:"shipping_address"`
	UserPhoneNumber string     `json:"user_phone_number" db:"user_phoneNumber"`
	State           string     `json:"state" db:"state"`
	Coupon          *uuid.UUID `json:"coupon_id" db:"coupon_id"`
	CouponDiscount  int        `json:"coupon_discount" db:"coupon_discount"`
	PaidAt          time.Time  `json:"paid_at" db:"paid_at"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`
}

type OrderProduct struct {
	ID           uuid.UUID `json:"id" db:"id"`
	OrderId      uuid.UUID `json:"order_id" db:"order_id"`
	VariantId    uuid.UUID `json:"variant_id" db:"variant_id"`
	VariantName  string    `json:"variant_name" db:"variant_name"`
	VariantPrice string    `json:"variant_price" db:"variant_price"`
	Quantity     int       `json:"quantity" db:"quantity"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type CreateCartOrder struct {
	CartId          uuid.UUID `json:"cart_id"`
	CouponId        uuid.UUID `json:"coupon_id"`
	ShippingAddress string    `json:"shipping_address"`
}

type CreateOrderProduct struct{}
