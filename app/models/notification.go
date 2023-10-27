package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type NotificationUser struct {
	ID             uuid.UUID `json:"id" db:"id"`
	UserId         uuid.UUID `json:"user_id" db:"user_id"`
	NotificationId uuid.UUID `json:"notification_id" db:"notification_id"`
	State          string    `json:"state" db:"state"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

type CreateNotification struct {
	Title   string `json:"title" db:"title"`
	Message string `json:"message" db:"message"`
}

type CreateUserNotication struct {
	UserId         uuid.UUID
	NotificationId uuid.UUID
	State          string
}
