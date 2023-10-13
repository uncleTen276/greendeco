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
	VariantTable               = "variants"
	DefaultProductVariantTable = "default_product_variant"
)

var _ VariantRepository = (*VariantRepo)(nil)

func NewVariantRepo(db *database.DB) VariantRepository {
	return &VariantRepo{db: db}
}

func (repo *VariantRepo) Create(m *models.CreateVariant) error {
	if m.IsDefault {
		if err := repo.createDefaultVariant(m); err != nil {
			return err
		}

		return nil
	}

	if err := repo.createVariant(m); err != nil {
		return err
	}

	return nil
}

func (repo *VariantRepo) createVariant(m *models.CreateVariant) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (product, name, color, price,currency, image, description, available) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`, VariantTable)
	_, err := repo.db.Exec(query, m.ProductId, m.Name, m.Color, m.Price, m.Currency, m.Image, m.Description, m.Available)
	if err != nil {
		return err
	}

	return nil
}

func (repo *VariantRepo) createDefaultVariant(m *models.CreateVariant) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := fmt.Sprintf(`INSERT INTO "%s" (product, name, color, price,currency, image, description, available) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`, VariantTable)
	newVariant := tx.QueryRow(query, m.ProductId, m.Name, m.Color, m.Price, m.Currency, m.Image, m.Description, m.Available)
	if err != nil {
		return err
	}

	var variantId string
	if err := newVariant.Scan(&variantId); err != nil {
		return err
	}

	createDefaultQuery := fmt.Sprintf(`INSERT INTO "%s" (product_id, variant_id) VALUES ($1,$2)`, DefaultProductVariantTable)
	_, err = tx.Exec(createDefaultQuery, m.ProductId, variantId)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
