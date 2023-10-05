package repository

import (
	"fmt"

	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type ProductRepository interface {
	CreateCategory(m *models.CreateCategory) error
	UpdateCategoryById(m *models.UpdateCategory) error
	FindCategoryById(id string) (*models.Category, error)
	DeleteCategory(id string) error
}

type ProductRepo struct {
	db *database.DB
}

const (
	CategoryTable       = "categories"
	CategoryNameField   = "name"
	CategoryParentField = "parent"
)

var _ ProductRepository = (*ProductRepo)(nil)

func NewProductRepository(db *database.DB) ProductRepository {
	return &ProductRepo{db}
}

func (repo *ProductRepo) CreateCategory(m *models.CreateCategory) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (name) VALUES ($1)`, CategoryTable)
	_, err := repo.db.Exec(query, m.Name)
	return err
}

func (repo *ProductRepo) UpdateCategoryById(m *models.UpdateCategory) error {
	query := fmt.Sprintf(`UPDATE "%s" SET name = $2 WHERE id = $1`, CategoryTable)
	_, err := repo.db.Exec(query, m.ID, m.Name)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepo) FindCategoryById(id string) (*models.Category, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, CategoryTable)
	category := models.NewCategory()
	err := repo.db.Get(category, query, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *ProductRepo) DeleteCategory(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, CategoryTable)
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *ProductRepo) GetCategories() ([]models.Category, error) {
	return nil, nil
}
