package repository

import (
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type UserRepository interface {
	Create(u models.CreateUser) error
}

type UserRepo struct {
	db *database.DB
}

var _ UserRepository = (*UserRepo)(nil)

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}

func (repo *UserRepo) Create(u models.CreateUser) error {
	return nil
}
