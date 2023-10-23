package repository

import (
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type OrderRepository interface {
	CreateOrder(*models.CreateOrder) (string, error)
	CreateOrderProduct(*models.CreateOrderProduct) error
}

type OrderRepo struct {
	db *database.DB
}

const (
	OrderTable        = "orders"
	OrderProductTable = "order_variants"
)

var _ OrderRepository = (*OrderRepo)(nil)

func NewOrderRepo(db *database.DB) OrderRepository {
	return &OrderRepo{
		db: db,
	}
}

func (repo *OrderRepo) CreateOrder(m *models.CreateOrder) (string, error) {
	return "", nil
}

func (repo *OrderRepo) CreateOrderProduct(m *models.CreateOrderProduct) error {
	return nil
}
