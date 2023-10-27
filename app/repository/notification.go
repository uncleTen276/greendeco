package repository

import (
	"fmt"

	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type NotificationRepository interface {
	Create(*models.CreateNotification) (string, error)
	CreateNotificationUser(*models.CreateUserNotication) error
}

type NotificationRepo struct {
	db *database.DB
}

const (
	NotificationTable = "notifications"
)

var _ NotificationRepository = (*NotificationRepo)(nil)

func NewNotificationRepo(db *database.DB) NotificationRepository {
	return &NotificationRepo{db: db}
}

func (repo *NotificationRepo) Create(m *models.CreateNotification) (string, error) {
	query := fmt.Sprintf(`INSERT INTO "%s" (title,message) VALUES ($1,$2) RETURNING id`, NotificationTable)
	newNoti := repo.db.QueryRow(query, m.Title, m.Message)
	var notiId string
	if err := newNoti.Scan(&notiId); err != nil {
		return "", err
	}

	return notiId, nil
}

func (repo *NotificationRepo) CreateNotificationUser(m *models.CreateUserNotication) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (user_id, notification_id, state) VALUES($1,$2,$3)`)
	if _, err := repo.db.Exec(query, m.UserId, m.NotificationId, m.State); err != nil {
		return err
	}

	return nil
}
