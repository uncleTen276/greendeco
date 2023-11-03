package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type CouponRepository interface {
	Create(*models.CreateCoupon) (string, error)
	Update(*models.UpdateCoupon) error
	GetCouponById(uuid.UUID) (*models.Coupon, error)
	GetCouponByCode(string) (*models.Coupon, error)
	DeleteById(id uuid.UUID) error
}

type CouponRepo struct {
	db *database.DB
}

const (
	CouponTable = "coupons"
)

var _ CouponRepository = (*CouponRepo)(nil)

func NewCouponRepo(db *database.DB) CouponRepository {
	return &CouponRepo{db: db}
}

func (repo *CouponRepo) Create(m *models.CreateCoupon) (string, error) {
	query := fmt.Sprintf(`INSERT INTO "%s" (name, discount, code, description,start_date, end_date) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`, CouponTable)
	newColor := repo.db.QueryRow(query, m.Name, m.Discount, m.Code, m.Description, m.StartDate, m.EndDate)
	var couponId string
	if err := newColor.Scan(&couponId); err != nil {
		return "", err
	}

	return couponId, nil
}

func (repo *CouponRepo) Update(m *models.UpdateCoupon) error {
	query := fmt.Sprintf(`UPDATE "%s" SET name=$2, code=$3, discord=$4, description=$5,start_date=$6, end_date=$7 WHERE id =$1 `, CouponTable)
	if _, err := repo.db.Exec(query, m.ID, m.Name, m.Code, m.Discount, m.Description, m.StartDate, m.EndDate); err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNotFound
		}
		return err
	}

	return nil
}

func (repo *CouponRepo) GetCouponById(id uuid.UUID) (*models.Coupon, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, CouponTable)
	coupon := &models.Coupon{}
	err := repo.db.Get(coupon, query, id)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return coupon, nil
}

func (repo *CouponRepo) GetCouponByCode(code string) (*models.Coupon, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE code = $1`, CouponTable)
	coupon := &models.Coupon{}
	err := repo.db.Get(coupon, query, code)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return coupon, nil
}

func (repo *CouponRepo) DeleteById(id uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, CouponTable)
	if _, err := repo.db.Exec(query, id); err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNotFound
		}

		return err
	}

	return nil
}
