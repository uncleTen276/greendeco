package repository

import (
	"database/sql"

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
}

type UserRepo struct {
	db *database.DB
}

var _ UserRepository = (*UserRepo)(nil)

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}

func (repo *UserRepo) Create(u *models.CreateUser) error {
	query := `INSERT INTO "users" (email,identifier,password,first_name,last_name, phone_number) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := repo.db.Exec(query, u.Email, u.Identifier, u.Password, u.FirstName, u.LastName, u.PhoneNumber)
	return err
}

func (repo *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	user := models.NewUser()
	query := `SELECT * FROM "users" WHERE email = $1`
	err := repo.db.Get(user, query, email)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepo) GetUserByIdentifier(identifier string) (*models.User, error) {
	user := models.NewUser()
	query := `SELECT * FROM "users" WHERE identifier = $1`
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
	query := `SELECT * FROM "users" WHERE phone_number = $1`
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
	query := `SELECT * FROM "users" WHERE id = $1`
	err := repo.db.Get(user, query, uId)
	return user, err
}

func (repo *UserRepo) UpdatePasswordById(password string, id string) error {
	query := `UPDATE "users" SET password = $1 WHERE id = $2`
	_, err := repo.db.Exec(query, password, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) UpdateUserInfor() error {
	return nil
}

func (repo *UserRepo) UpdateRules() error {
	return nil
}
