package repository

import (
	"fmt"

	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type VariantRepository interface {
	Create(m *models.CreateVariant) error
}

type VariantRepo struct {
	db *database.DB
}

const (
	VariantTable = "variants"
)

var _ VariantRepository = (*VariantRepo)(nil)

func NewVariantRepo(db *database.DB) VariantRepository {
	return &VariantRepo{db: db}
}

func (repo *VariantRepo) Create(m *models.CreateVariant) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (is_default, product, name, color, price,currency, image, description) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`, VariantTable)
	_, err := repo.db.Exec(query, m.IsDefault, m.ProductId, m.Name, m.Color, m.Price, m.Currency, m.Image, m.Description)
	if err != nil {
		return err
	}

	return nil
}
