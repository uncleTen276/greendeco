package models

import (
	"time"
)

type User struct {
	ID          string    `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Email       string    `db:"email"`
	Identifier  string    `db:"identifier"`
	Password    string    `db:"password" json:"-"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	PhoneNumber string    `db:"phone_number"`
	IsAdmin     bool      `db:"admin"`
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

type UserTokens struct {
	AccessToken  string `json:"access_Token"`
	RefreshToken string `json:"refresh_Token"`
}

type EmailResponse struct {
	Link     string
	Username string
}
