package repository

import (
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type ProductRepository interface {
	Create(m *models.CreateProduct) error
	// UpdateById(m *models.UpdateP) error
	// FindById(id string) (*models.Category, error)
	// Delete(id string) error
	// All(limit, offset int) ([]*models.Category, error)
}

type ProductRepo struct {
	db *database.DB
}

var _ ProductRepository = (*ProductRepo)(nil)

func (repo *ProductRepo) Create(m *models.CreateProduct) error {
	return nil
}

func UpdateById(m *models.UpdateCategory) error {
	return nil
}
