package repository

import (
	"fmt"

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

const (
	ProductTable = "products"
)

var _ ProductRepository = (*ProductRepo)(nil)

func NewProductRepository(db *database.DB) ProductRepository {
	return &ProductRepo{db: db}
}

func (repo *ProductRepo) Create(m *models.CreateProduct) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (category_id ,name, images, size, type, detail, light, difficulty, warter, qr_image) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`, ProductTable)
	_, err := repo.db.Exec(query, m.CategoryId, m.Name, m.Images, m.Size, m.Type, m.Detail, m.Light, m.Difficulty, m.Warter, m.QrImage)
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func UpdateById(m *models.UpdateCategory) error {
	return nil
}
