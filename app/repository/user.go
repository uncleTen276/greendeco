package repository

import (
	"database/sql"
	"fmt"

	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type UserRepository interface {
	Create(u *models.CreateUser) error
	GetUserByIdentifier(identifier string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByPhoneNumber(phoneNumber string) (*models.User, error)
	GetUserById(uId string) (*models.User, error)
	UpdatePasswordById(password string, id string) error
	UpdateUserInfor(userId string, user *models.UpdateUser) error
}

type UserRepo struct {
	db *database.DB
}

const (
	UserTable = "users"
)

var _ UserRepository = (*UserRepo)(nil)

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}

func (repo *UserRepo) Create(u *models.CreateUser) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (email,identifier,password,first_name,last_name, phone_number) VALUES ($1,$2,$3,$4,$5,$6)`, UserTable)
	_, err := repo.db.Exec(query, u.Email, u.Identifier, u.Password, u.FirstName, u.LastName, u.PhoneNumber)
	return err
}

func (repo *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE email = $1`, UserTable)
	err := repo.db.Get(user, query, email)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepo) GetUserByIdentifier(identifier string) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE identifier = $1`, UserTable)
	err := repo.db.Get(user, query, identifier)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE phone_number = $1`, UserTable)
	err := repo.db.Get(user, query, phoneNumber)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) GetUserById(uId string) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, UserTable)
	err := repo.db.Get(user, query, uId)
	return user, err
}

func (repo *UserRepo) UpdatePasswordById(password string, id string) error {
	query := fmt.Sprintf(`UPDATE "%s" SET password = $1 WHERE id = $2`, UserTable)
	_, err := repo.db.Exec(query, password, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) UpdateUserInfor(userId string, user *models.UpdateUser) error {
	query := fmt.Sprintf(`UPDATE "%s" SET first_name = $2, last_name = $3, avatar = $4, phone_number = $5   WHERE id = $1`, UserTable)
	_, err := repo.db.Exec(query, userId, user.FirstName, user.LastName, user.Avatar, user.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) UpdateRules() error {
	return nil
}
