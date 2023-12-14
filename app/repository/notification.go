package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type NotificationRepository interface {
	Create(*models.CreateNotification) (string, error)
	CreateNotificationUser(*models.CreateUserNotication) error
	SendNotificationToUsers(*models.UserListNotification) error
	GetNotificationsByUserId(userId uuid.UUID, q *models.BaseQuery) ([]models.UserNotificationResponse, error)
	UpdateReadNotification(notiId uuid.UUID) error
	GetUserNotificationById(notiId uuid.UUID) (*models.NotificationUser, error)
	GetNotificationById(notiId uuid.UUID) (*models.Notification, error)
	UpdateNotificaionById(m *models.UpdateNotification) error
}

type NotificationRepo struct {
	db *database.DB
}

const (
	NotificationTable     = "notifications"
	NotificationUserTable = "notifications_users"
)

var _ NotificationRepository = (*NotificationRepo)(nil)

func NewNotificationRepo(db *database.DB) NotificationRepository {
	return &NotificationRepo{db: db}
}

func (repo *NotificationRepo) Create(m *models.CreateNotification) (string, error) {
	query := fmt.Sprintf(`INSERT INTO "%s" (title,message,description) VALUES ($1,$2,$3) RETURNING id`, NotificationTable)
	newNoti := repo.db.QueryRow(query, m.Title, m.Message, m.Description)
	var notiId string
	if err := newNoti.Scan(&notiId); err != nil {
		return "", err
	}

	return notiId, nil
}

func (repo *NotificationRepo) CreateNotificationUser(m *models.CreateUserNotication) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (user_id, notification_id, state) VALUES($1,$2,$3)`, NotificationUserTable)
	if _, err := repo.db.Exec(query, m.UserId, m.NotificationId, m.State); err != nil {
		return err
	}

	return nil
}

func (repo *NotificationRepo) SendNotificationToUsers(m *models.UserListNotification) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (user_id, notification_id, state) VALUES($1,$2,$3)`, NotificationUserTable)
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, v := range m.Users {
		if _, err := tx.Exec(query, v, m.Notification_Id, models.NotificationUnreadState); err != nil {
			if err == sql.ErrNoRows {
				return models.ErrNotFound
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo *NotificationRepo) SendUserNotification(notification *models.CreateNotification, user *models.CreateUserNotication) (string, error) {
	createNotiQuery := fmt.Sprintf(`INSERT INTO "%s" (title,message) VALUES ($1,$2) RETURNING id`, NotificationTable)
	tx, err := repo.db.Begin()
	if err != nil {
		return "", err
	}

	defer tx.Rollback()

	createNotiUserQuery := fmt.Sprintf(`INSERT INTO "%s" (user_id, notification_id, state) VALUES($1,$2,$3)`, NotificationUserTable)

	newNoti := tx.QueryRow(createNotiQuery, notification.Title, notification.Message)
	var newNotiId string

	err = newNoti.Scan(&newNotiId)
	if err != nil {
		return "", err
	}

	newUserNoti := tx.QueryRow(createNotiUserQuery, user.UserId, user.NotificationId, user.State)
	var userNotiId string
	err = newUserNoti.Scan(&userNotiId)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return userNotiId, nil
}

func (repo *NotificationRepo) GetNotificationsByUserId(userId uuid.UUID, q *models.BaseQuery) ([]models.UserNotificationResponse, error) {
	baseQuery := fmt.Sprintf(`SELECT nu.id, nu.notification_id, nu.state, nu.created_at, n.title, n.message, n.description FROM "%s" AS nu LEFT JOIN "%s" AS n ON n.id = nu.notification_id WHERE nu.user_id = $1 `, NotificationUserTable, NotificationTable)
	notifications := []models.UserNotificationResponse{}
	query := newQueryBuilder(baseQuery, "n.created_at").
		SortBy(q.SortBy, q.Sort).
		Build()

	limit := q.Limit
	limit += 1
	pageOffset := q.Limit * (q.OffSet - 1)
	query = fmt.Sprintf(query+" LIMIT %d OFFSET %d", limit, pageOffset)

	if err := repo.db.Select(&notifications, query, userId); err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}

		return nil, err
	}

	return notifications, nil
}

func (repo *NotificationRepo) UpdateReadNotification(notiId uuid.UUID) error {
	query := fmt.Sprintf(`UPDATE "%s" SET state='read' WHERE id=$1`, NotificationUserTable)
	if _, err := repo.db.Exec(query, notiId); err != nil {
		return err
	}

	return nil
}

func (repo *NotificationRepo) GetNotificationById(notiId uuid.UUID) (*models.Notification, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, NotificationTable)
	noti := &models.Notification{}
	if err := repo.db.Get(noti, query, notiId); err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}

		return nil, err
	}

	return noti, nil
}

func (repo *NotificationRepo) GetUserNotificationById(notiId uuid.UUID) (*models.NotificationUser, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, NotificationUserTable)
	noti := &models.NotificationUser{}
	if err := repo.db.Get(noti, query, notiId); err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}

		return nil, err
	}

	return noti, nil
}

func (repo *NotificationRepo) UpdateNotificaionById(m *models.UpdateNotification) error {
	query := fmt.Sprintf(`UPDATE "%s" SET title=$2, message=$3 WHERE id=$1`, NotificationTable)
	if _, err := repo.db.Exec(query, m.ID, m.Title, m.Message); err != nil {
		return err
	}

	return nil
}
