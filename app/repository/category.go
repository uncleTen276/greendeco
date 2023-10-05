package repository

import (
	"fmt"

	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type CategoryRepository interface {
	Create(m *models.CreateCategory) error
	UpdateById(m *models.UpdateCategory) error
	FindById(id string) (*models.Category, error)
	Delete(id string) error
	All(limit, offset int) ([]*models.Category, error)
}

type ProductRepo struct {
	db *database.DB
}

const (
	CategoryTable       = "categories"
	CategoryNameField   = "name"
	CategoryParentField = "parent"
)

var _ CategoryRepository = (*ProductRepo)(nil)

func NewCategoryRepository(db *database.DB) CategoryRepository {
	return &ProductRepo{db}
}

func (repo *ProductRepo) Create(m *models.CreateCategory) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (name) VALUES ($1)`, CategoryTable)
	_, err := repo.db.Exec(query, m.Name)
	return err
}

func (repo *ProductRepo) UpdateById(m *models.UpdateCategory) error {
	query := fmt.Sprintf(`UPDATE "%s" SET name = $2 WHERE id = $1`, CategoryTable)
	_, err := repo.db.Exec(query, m.ID, m.Name)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepo) FindById(id string) (*models.Category, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, CategoryTable)
	category := models.NewCategory()
	err := repo.db.Get(category, query, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *ProductRepo) Delete(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, CategoryTable)
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *ProductRepo) All(limit, offset int) ([]*models.Category, error) {
	categories := []*models.Category{}
	query := `SELECT * FROM "categories" LIMIT $1 OFFSET $2`
	if err := repo.db.Select(&categories, query, limit, offset); err != nil {
		return nil, err
	}

	return categories, nil
}
