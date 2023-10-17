package models

import (
	"time"
)

type User struct {
	ID          string    `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
	UpdatedAt   time.Time `db:"updated_at" json:"-"`
	Email       string    `db:"email" json:"email"`
	Identifier  string    `db:"identifier" json:"identifier"`
	Password    string    `db:"password" json:"-"`
	FirstName   string    `db:"first_name" json:"firstName"`
	LastName    string    `db:"last_name" json:"lastName"`
	PhoneNumber string    `db:"phone_number" json:"phoneNumber"`
	Avatar      *string   `db:"avatar" json:"avatar"`
	IsAdmin     bool      `db:"admin" json:"-"`
}

func NewUser() *User {
	return &User{}
}

type CreateUser struct {
	Email       string `json:"email" validate:"required,email,lte=150"`
	Identifier  string `json:"identifier" validate:"required,lte=100"`
	Password    string `json:"password" validate:"required,lte=50,gte=8"`
	FirstName   string `json:"firstName" validate:"required,lte=50"`
	LastName    string `json:"lastName" validate:"required,lte=50"`
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"`
}

type UserLogin struct {
	Identifier string `json:"identifier" validate:"required,lte=100"`
	Password   string `json:"password" validate:"required,lte=50,gte=8"`
}

type UserToken struct {
	AccessToken string `json:"access_Token"`
}

type EmailResponse struct {
	Link     string
	Username string
}

type UpdateUser struct {
	FirstName   string  `json:"firstName" validate:"required,lte=50"`
	LastName    string  `json:"lastName" validate:"required,lte=50"`
	PhoneNumber string  `json:"phoneNumber" validate:"required,e164"`
	Email       string  `json:"email" validate:"required,email,lte=150"`
	Avatar      *string `json:"avatar" validate:"omitempty,url"`
}
