package models

import (
	"time"

	"github.com/google/uuid"
)

const NotificationUnreadState = "unread"

type Notification struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Message     string    `json:"message" db:"message"`
	Description *string   `json:"description" db:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
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
	Title       string  `json:"title" db:"title" validate:"required,gte=3"`
	Message     string  `json:"message" db:"message" validate:"required,gte=3"`
	Description *string ` json:"description" db:"description"`
}

type CreateUserNotication struct {
	UserId         uuid.UUID
	NotificationId uuid.UUID
	State          string
	Description    string
}

type UserListNotification struct {
	Users           []uuid.UUID `json:"users"`
	Notification_Id *uuid.UUID  `json:"notification_id"`
}

type UserNotificationResponse struct {
	ID             uuid.UUID `json:"id" db:"id"`
	NotificationId uuid.UUID `json:"notification_id" db:"notification_id"`
	State          string    `json:"state" db:"state"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	Title          string    `json:"title" db:"title"`
	Message        string    `json:"message" db:"message"`
	Description    *string   `json:"description" db:"description"`
}

type UpdateNotification struct {
	ID      uuid.UUID `json:"-" db:"id" validate:"required,uuid4"`
	Title   string    `json:"title" db:"title" validate:"required,gte=3"`
	Message string    `json:"message" db:"message" validate:"required,gte=3"`
}
