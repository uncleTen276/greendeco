package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type VariantRepository interface {
	Create(m *models.CreateVariant) error
	FindById(id uuid.UUID) (*models.Variant, error)
	UpdateById(m *models.UpdateVariant) error
	GetVariantsByProductId(id uuid.UUID) ([]models.Variant, error)
	GetDefaultVariantOfProduct(id uuid.UUID) (*models.DefaultVariant, error)
	UpdateDefaultVariant(m *models.UpdateDefaultVariant) error
	Delete(id uuid.UUID) error
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
	query := fmt.Sprintf(`INSERT INTO "%s" (product, name, color, color_name, price,currency, image, description, available) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`, VariantTable)
	_, err := repo.db.Exec(query, m.ProductId, m.Name, m.Color, m.ColorName, m.Price, m.Currency, m.Image, m.Description, m.Available)
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

	query := fmt.Sprintf(`INSERT INTO "%s" (product, name, color,color_name, price,currency, image, description, available) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`, VariantTable)
	newVariant := tx.QueryRow(query, m.ProductId, m.Name, m.Color, m.ColorName, m.Price, m.Currency, m.Image, m.Description, m.Available)
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

func (repo *VariantRepo) GetVariantsByProductId(id uuid.UUID) ([]models.Variant, error) {
	result := []models.Variant{}
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE product = $1`, VariantTable)
	if err := repo.db.Select(&result, query, id); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *VariantRepo) FindById(id uuid.UUID) (*models.Variant, error) {
	variant := &models.Variant{}
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, VariantTable)
	err := repo.db.Get(variant, query, id)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return variant, nil
}

func (repo *VariantRepo) Delete(id uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, VariantTable)
	if _, err := repo.db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (repo *VariantRepo) GetDefaultVariantOfProduct(id uuid.UUID) (*models.DefaultVariant, error) {
	query := fmt.Sprintf(`SELECT product_id, variant_id FROM "%s" WHERE product_id  = $1`, DefaultProductVariantTable)
	defaultVariant := &models.DefaultVariant{}
	if err := repo.db.Select(defaultVariant, query, id); err != nil {
		return nil, err
	}

	return defaultVariant, nil
}

func (repo *VariantRepo) UpdateDefaultVariant(m *models.UpdateDefaultVariant) error {
	query := fmt.Sprintf(`UPDATE %s SET variant_id = $1 WHERE product_id = $2`, ProductVariantDefaultTable)
	if _, err := repo.db.Exec(query, m.VariantId, m.ProductId); err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func (repo *VariantRepo) UpdateById(m *models.UpdateVariant) error {
	query := fmt.Sprintf(`UPDATE "%s" SET available = $2, name = $3, color = $4, price = $5, currency = $6, image = $7, description = $8, color_name = $9  WHERE id = $1`, VariantTable)
	if _, err := repo.db.Exec(query, m.ID, m.Available, m.Name, m.Color, m.Price, m.Currency, m.Image, m.Description, m.ColorName); err != nil {
		println(err.Error())
		return err
	}
	return nil
}
