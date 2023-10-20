package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type CartRepository interface {
	Create(*models.CreateCart) (string, error)
	CreateCartProduct(*models.CreateCartProduct) (string, error)
	GetCartByOwnerId(uuid.UUID) (*models.Cart, error)
}

type CartRepo struct {
	db *database.DB
}

const (
	CartTable        = "carts"
	CartProductTable = "cart_variants"
)

var _ CartRepository = (*CartRepo)(nil)

func NewCartRepo(db *database.DB) CartRepository {
	return &CartRepo{db: db}
}

func (repo *CartRepo) Create(m *models.CreateCart) (string, error) {
	query := fmt.Sprintf(`INSERT INTO "%s" (owner_id, description) VALUES ($1,$2) RETURNING id`, CartTable)
	newCart := repo.db.QueryRow(query, m.Owner, m.Description)
	var cartId string
	if err := newCart.Scan(&cartId); err != nil {
		return "", err
	}

	return cartId, nil
}

func (repo *CartRepo) CreateCartProduct(m *models.CreateCartProduct) (string, error) {
	query := fmt.Sprintf(`INSERT INTO "%s" (cart_id, variant_id, quantity) VALUES ($1,$2, $3) RETURNING id`, CartProductTable)
	newCartItem := repo.db.QueryRow(query, m.Cart, m.Variant, m.Quantity)
	var cartItemId string
	if err := newCartItem.Scan(&cartItemId); err != nil {
		if err == sql.ErrNoRows {
			return "", models.ErrNotFound
		}

		return "", err
	}

	return cartItemId, nil
}

func (repo *CartRepo) GetCartByOwnerId(ownerId uuid.UUID) (*models.Cart, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE owner_id = $1`, CartTable)
	cart := &models.Cart{}
	err := repo.db.Get(cart, query, ownerId)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return cart, nil
}
