package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type ProductRepository interface {
	Create(m *models.CreateProduct) (string, error)
	UpdateById(m *models.UpdateProduct) error
	FindById(id uuid.UUID) (*models.Product, error)
	Delete(id uuid.UUID) error
	All(q models.ProductQuery) ([]models.ActivedProduct, error)
	GetRecommendProducts(id uuid.UUID) ([]string, error)
	CreateRecommendProduct(m *models.CreateRecommend) error
	DeleteRecommendProduct(m *models.CreateRecommend) error
	GetAllProducts(q *models.ProductQuery) ([]models.Product, error)
}
type ProductRepo struct {
	db *database.DB
}

const (
	ProductTable               = "products"
	RecommendTable             = "recommends"
	ProductVariantView         = "published_products"
	ProductVariantDefaultTable = "default_product_variant"
)

var _ ProductRepository = (*ProductRepo)(nil)

func NewProductRepo(db *database.DB) ProductRepository {
	return &ProductRepo{db: db}
}

func (repo *ProductRepo) Create(m *models.CreateProduct) (string, error) {
	query := fmt.Sprintf(`INSERT INTO "%s" (category_id ,name, images, size, type, detail, light, difficulty, water, description ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9, $10) RETURNING id`, ProductTable)
	newProduct := repo.db.QueryRow(query, m.CategoryId, m.Name, m.Images, m.Size, m.Type, m.Detail, m.Light, m.Difficulty, m.Water, m.Description)
	var productId string
	if err := newProduct.Scan(&productId); err != nil {
		return "", err
	}

	return productId, nil
}

func (repo *ProductRepo) UpdateById(m *models.UpdateProduct) error {
	query := fmt.Sprintf(`UPDATE "%s" SET is_publish = $2, size = $3, type=$4,images = $5, description = $6, detail = $7, light = $8, difficulty = $9, water = $10, available = $11  WHERE id = $1`, ProductTable)
	if _, err := repo.db.Exec(query, m.ID, m.IsPublish, m.Size, m.Type, m.Images, m.Description, m.Detail, m.Light, m.Difficulty, m.Water, m.Available); err != nil {
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
	limit += 1
	pageOffset := q.BaseQuery.Limit * (q.BaseQuery.OffSet - 1)

	results := []models.ActivedProduct{}
	firstQuery := fmt.Sprintf(`SELECT * FROM "%s" `, ProductVariantView)
	query := repo.newProductQueryBuilder(firstQuery).
		SetName(q.Fields.Name).
		SetAvailable(q.Fields.Available).
		SetCategory(q.Fields.Category).
		SetSize(q.Fields.Size).
		SetType(q.Fields.Type).
		SetDifficulty(q.Fields.Difficulty).
		SetWater(q.Fields.Water).
		SortBy(q.SortBy, q.Sort).
		Build()

	query = fmt.Sprintf(query+" LIMIT %d OFFSET %d", limit, pageOffset)
	if err := repo.db.Select(&results, query); err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *ProductRepo) FindById(id uuid.UUID) (*models.Product, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, ProductTable)
	product := &models.Product{}
	err := repo.db.Get(product, query, id)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *ProductRepo) GetRecommendProducts(id uuid.UUID) ([]string, error) {
	results := []string{}
	query := fmt.Sprintf(`SELECT recommend_product FROM "%s" WHERE product_id = $1`, RecommendTable)
	err := repo.db.Select(&results, query, id)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *ProductRepo) CreateRecommendProduct(m *models.CreateRecommend) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (product_id, recommend_product) VALUES ($1,$2)`, RecommendTable)
	if _, err := repo.db.Exec(query, m.ProductId, m.RecommendId); err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepo) DeleteRecommendProduct(m *models.CreateRecommend) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE product = $1 AND recommend = $2`, RecommendTable)
	if _, err := repo.db.Exec(query, m.ProductId, m.RecommendId); err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepo) DeleteCategory(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, CategoryTable)
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *ProductRepo) GetAllProducts(q *models.ProductQuery) ([]models.Product, error) {
	limit := q.Limit
	limit += 1
	pageOffset := q.BaseQuery.Limit * (q.BaseQuery.OffSet - 1)

	results := []models.Product{}
	firstQuery := fmt.Sprintf(`SELECT * FROM "%s" `, ProductTable)
	query := repo.newProductQueryBuilder(firstQuery).
		SetName(q.Fields.Name).
		SetAvailable(q.Fields.Available).
		SetCategory(q.Fields.Category).
		SetSize(q.Fields.Size).
		SetType(q.Fields.Type).
		SetDifficulty(q.Fields.Difficulty).
		SetWater(q.Fields.Water).
		SetPublished(q.Fields.IsPublish).
		SortBy(q.SortBy, q.Sort).
		Build()

	query = fmt.Sprintf(query+" LIMIT %d OFFSET %d", limit, pageOffset)
	if err := repo.db.Select(&results, query); err != nil {
		return nil, err
	}

	return results, nil
}
