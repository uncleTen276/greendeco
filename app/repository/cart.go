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
	GetCartProductById(cartId uuid.UUID) (*models.CartProduct, error)
	UpdateCartProductById(updateCart *models.UpdateCartProduct) error
	DeleteCartById(cartId uuid.UUID) error
	DeleteCartItemByCartId(cartId uuid.UUID) error
	DeleteCartItemById(itemId uuid.UUID) error
	GetCartById(cartId uuid.UUID) (*models.Cart, error)
	GetCartProductByCartId(cartId uuid.UUID, query *models.BaseQuery) ([]*models.CartProduct, error)
	GetAllCartProductByCartId(cartId uuid.UUID) ([]*models.CartProduct, error)
	DeleteCartItemByVariantId(variantId uuid.UUID) error
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
	query := fmt.Sprintf(`
INSERT INTO "%s" (cart_id, variant_id, quantity) VALUES ($1,$2, $3) ON CONFLICT (cart_id,variant_id) DO UPDATE SET quantity=%s.quantity+1  RETURNING id`, CartProductTable, CartProductTable)
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

func (repo *CartRepo) GetCartProductById(cartId uuid.UUID) (*models.CartProduct, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, CartProductTable)
	cartItem := &models.CartProduct{}
	err := repo.db.Get(cartItem, query, cartId)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (repo *CartRepo) UpdateCartProductById(updateCart *models.UpdateCartProduct) error {
	query := fmt.Sprintf(`UPDATE "%s" SET quantity=$2 WHERE id=$1`, CartProductTable)
	if _, err := repo.db.Exec(query, updateCart.ID, updateCart.Quantity); err != nil {
		return err
	}

	return nil
}

func (repo *CartRepo) DeleteCartById(cartId uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, CartTable)
	if _, err := repo.db.Exec(query, cartId); err != nil {
		return err
	}

	return nil
}

func (repo *CartRepo) DeleteCartItemByCartId(cartId uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE cart_id= $1`, CartProductTable)
	if _, err := repo.db.Exec(query, cartId); err != nil {
		return err
	}

	return nil
}

func (repo *CartRepo) DeleteCartItemById(itemId uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, CartProductTable)
	if _, err := repo.db.Exec(query, itemId); err != nil {
		return err
	}

	return nil
}

func (repo *CartRepo) GetCartById(cartId uuid.UUID) (*models.Cart, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, CartTable)
	cart := &models.Cart{}
	err := repo.db.Get(cart, query, cartId)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return cart, nil
}

func (repo *CartRepo) GetCartProductByCartId(cartId uuid.UUID, q *models.BaseQuery) ([]*models.CartProduct, error) {
	limit := q.Limit
	limit += 1
	pageOffset := q.Limit * (q.OffSet - 1)
	firstQuery := fmt.Sprintf(`SELECT * FROM "%s" WHERE cart_id = $1 `, CartProductTable)
	cartProductList := &[]*models.CartProduct{}
	query := repo.newCartQueryBuilder(firstQuery).
		SortBy(q.SortBy, q.Sort).
		Build()

	query = fmt.Sprintf(query+" LIMIT %d OFFSET %d", limit, pageOffset)
	err := repo.db.Select(cartProductList, query, cartId)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return *cartProductList, nil
}

func (repo *CartRepo) GetAllCartProductByCartId(cartId uuid.UUID) ([]*models.CartProduct, error) {
	cartProductList := []*models.CartProduct{}
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE cart_id = $1`, CartProductTable)
	err := repo.db.Select(&cartProductList, query, cartId)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return cartProductList, nil
}

func (repo *CartRepo) DeleteCartItemByVariantId(variantId uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE variant_id = $1`, CartProductTable)
	_, err := repo.db.Exec(query, variantId)
	if err == sql.ErrNoRows {
		return models.ErrNotFound
	} else if err != nil {
		return err
	}
	return nil
}
