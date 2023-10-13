package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type ProductRepository interface {
	Create(m *models.CreateProduct) error
	UpdateById(m *models.UpdateProduct) error
	// FindById(id string) (*models.Category, error)
	Delete(id uuid.UUID) error
	All(q models.ProductQuery) ([]models.ActivedProduct, error)
}

type ProductRepo struct {
	db *database.DB
}

const (
	ProductTable       = "products"
	ProductVariantView = "published_products"
)

var _ ProductRepository = (*ProductRepo)(nil)

func NewProductRepo(db *database.DB) ProductRepository {
	return &ProductRepo{db: db}
}

func (repo *ProductRepo) Create(m *models.CreateProduct) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (category_id ,name, images, size, type, detail, light, difficulty, warter ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`, ProductTable)
	_, err := repo.db.Exec(query, m.CategoryId, m.Name, m.Images, m.Size, m.Type, m.Detail, m.Light, m.Difficulty, m.Warter)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepo) UpdateById(m *models.UpdateProduct) error {
	query := fmt.Sprintf(`UPDATE "%s" SET is_publish = $2, size = $3, type=$4,images = $5, description = $6, detail = $7, light = $8, difficulty = $9, warter = $10, available = $11  WHERE id = $1`, ProductTable)
	if _, err := repo.db.Exec(query, m.ID, m.IsPublish, m.Size, m.Type, m.Images, m.Description, m.Detail, m.Light, m.Difficulty, m.Warter, m.Available); err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepo) Delete(id uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, ProductTable)
	if _, err := repo.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepo) All(q models.ProductQuery) ([]models.ActivedProduct, error) {
	limit := q.Limit
	if !q.BaseQuery.IsFirstPage() {
		limit += 1
	}
	pageOffset := q.BaseQuery.Limit * (q.BaseQuery.OffSet - 1)

	results := []models.ActivedProduct{}
	firstQuery := fmt.Sprintf(`SELECT * FROM "%s" `, ProductVariantView)
	query := NewProductQueryBuilder(firstQuery).
		SetName(q.Fields.Name).
		SetAvailable(q.Fields.Available).
		SetCategory(q.Fields.Category).
		SetSize(q.Fields.Size).
		SetType(q.Fields.Type).
		SetDifficulty(q.Fields.Difficulty).
		SetWarter(q.Fields.Warter).
		SortBy(q.SortBy, q.Sort).
		Build()

	println(q.Sort)
	query = fmt.Sprintf(query+" LIMIT %d OFFSET %d", limit, pageOffset)
	println(query)
	if err := repo.db.Select(&results, query); err != nil {
		return nil, err
	}

	return results, nil
}
