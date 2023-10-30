package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type OrderRepository interface {
	CreateOrderProduct(*models.CreateOrderProduct) error
	CreateOrderFromCart(m *models.Order, orderItem []*models.OrderProduct, cartId uuid.UUID) (string, error)
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

func (repo *OrderRepo) CreateOrderProduct(m *models.CreateOrderProduct) error {
	return nil
}

func (repo *OrderRepo) CreateOrderFromCart(m *models.Order, orderItems []*models.OrderProduct, cartId uuid.UUID) (string, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	orderQuery := fmt.Sprintf(`INSERT INTO "%s" (owner_id, user_name, user_email, user_phoneNumber, shipping_address, state, coupon_id, coupon_discount) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`, OrderTable)
	orderProductQuery := fmt.Sprintf(`INSERT INTO "%s" (order_id, variant_id, variant_name, variant_price, quantity) VALUES ($1,$2,$3,$4,$5)`, OrderProductTable)
	cartQuery := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, CartTable)

	newOrder := tx.QueryRow(orderQuery, m.OwnerId, m.UserName, m.UserEmail, m.UserPhoneNumber, m.ShippingAddress, m.State, m.Coupon, m.CouponDiscount)
	var newOrderId string
	err = newOrder.Scan(&newOrderId)
	if err != nil {
		return "", err
	}

	for _, item := range orderItems {
		_, err = tx.Exec(orderProductQuery, newOrderId, item.VariantId, item.VariantName, item.VariantPrice, item.Quantity)
		if err != nil {
			return "", err
		}
	}

	if _, err := tx.Exec(cartQuery, cartId); err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return newOrderId, nil
}
