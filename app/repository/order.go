package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type OrderRepository interface {
	CreateOrderFromCart(m *models.Order, orderItem []*models.OrderProduct, cartId uuid.UUID) (string, error)
	GetOrderById(orderId uuid.UUID) (*models.Order, error)
	GetOrderProductsByOrderId(orderId uuid.UUID, q *models.BaseQuery) ([]*models.OrderProductResponse, error)
	GetTotalPaymentForOrder(orderId uuid.UUID) (int, error)
	UpdateOrder(m *models.UpdateOrder) error
	All(q *models.OrderQuery) ([]*models.Order, error)
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

func (repo *OrderRepo) GetOrderById(orderId uuid.UUID) (*models.Order, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, OrderTable)
	order := &models.Order{}
	if err := repo.db.Get(order, query, orderId); err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}

		return nil, err
	}

	return order, nil
}

func (repo *OrderRepo) GetOrderProductsByOrderId(orderId uuid.UUID, q *models.BaseQuery) ([]*models.OrderProductResponse, error) {
	limit := q.Limit
	limit += 1
	pageOffset := q.Limit * (q.OffSet - 1)
	firstQuery := fmt.Sprintf(`SELECT t1.* , t2.product FROM "%s" AS t1 LEFT JOIN "%s" AS t2 ON t1.variant_id = t2.id WHERE order_id = $1 `, OrderProductTable, VariantTable)
	query := repo.newOrderQueryBuilder(firstQuery).
		SortBy(q.SortBy, q.Sort).
		Build()

	query = fmt.Sprintf(query+" LIMIT %d OFFSET %d", limit, pageOffset)
	products := []*models.OrderProductResponse{}
	if err := repo.db.Select(&products, query, orderId); err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}

		return nil, err
	}

	return products, nil
}

func (repo *OrderRepo) UpdateOrder(m *models.UpdateOrder) error {
	query := fmt.Sprintf(`UPDATE "%s" SET state = $2, paid_at = $3  WHERE id = $1`, OrderTable)
	if _, err := repo.db.Exec(query, m.OrderId, m.State, m.PaidAt); err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNotFound
		}

		return err
	}

	return nil
}

func (repo *OrderRepo) GetTotalPaymentForOrder(orderId uuid.UUID) (int, error) {
	query := fmt.Sprintf(`SELECT SUM(variant_price * quantity) FROM "%s" WHERE order_id = $1`, OrderProductTable)
	result := repo.db.QueryRow(query, orderId)
	var total int
	if err := result.Scan(&total); err != nil {
		if err == sql.ErrNoRows {
			return 0, models.ErrNotFound
		}
		return 0, err
	}

	return total, nil
}

func (repo *OrderRepo) All(q *models.OrderQuery) ([]*models.Order, error) {
	limit := q.Limit
	limit += 1
	pageOffset := q.Limit * (q.OffSet - 1)
	firstQuery := fmt.Sprintf(`SELECT * FROM "%s" `, OrderTable)
	query := repo.newOrderQueryBuilder(firstQuery).
		SetOwner(q.Fields.OwnerId).
		SetState(q.Fields.State).
		SetCoupon(q.Fields.Coupon).
		SortBy(q.SortBy, q.Sort).
		Build()

	query = fmt.Sprintf(query+" LIMIT %d OFFSET %d", limit, pageOffset)

	orders := []*models.Order{}
	if err := repo.db.Select(&orders, query); err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}

		return nil, err
	}

	return orders, nil
}
