package models

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Discount    int       `db:"discount" json:"discount"`
	Code        string    `db:"code" json:"code"`
	Description string    `db:"description" json:"description"`
	StartDate   time.Time `db:"start_date" json:"start_date"`
	EndDate     time.Time `db:"end_date" json:"end_date"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type CreateCoupon struct {
	Name        string `db:"name" json:"name" validate:"required"`
	Discount    int    `db:"discount" json:"discount" validate:"required,numeric"`
	Code        string `db:"code" json:"code" validate:"required,lte=5"`
	Description string `db:"description" json:"description"`
	StartDate   string `db:"start_date" json:"start_date" `
	EndDate     string `db:"end_date" json:"end_date"`
}

type UpdateCoupon struct {
	ID          uuid.UUID `db:"id" json:"-" validate:"required,uuid4"`
	Name        string    `db:"name" json:"name" validate:"required"`
	Discount    int       `db:"discount" json:"discount" validate:"required,numeric"`
	Code        string    `db:"code" json:"code" validate:"required,lte=5"`
	Description string    `db:"description" json:"description"`
	StartDate   time.Time `db:"start_date" json:"start_date"`
	EndDate     time.Time `db:"end_date" json:"end_date"`
}
