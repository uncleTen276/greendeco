package models

import (
	"time"
)

type User struct {
	ID          int       `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Username    string    `db:"username"`
	Email       string    `db:"email"`
	Identifier  string    `db:"identifier"`
	Password    string    `db:"password"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	BirthDate   time.Time `db:"birthdate"`
	PhoneNumber string    `db:"phone_number"`
}

func NewUser() *User {
	return &User{}
}

type CreateUser struct {
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Identifier  string    `json:"identifier"`
	Password    string    `json:"password"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	BirthDate   time.Time `json:"birthdate"`
	PhoneNumber string    `json:"phoneNumber"`
}
