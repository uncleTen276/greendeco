package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type ColorRepository interface {
	Create(*models.CreateColor) (string, error)
	GetColorById(uuid.UUID) (*models.Color, error)
	UpdateColor(*models.UpdateColor) error
	All() ([]*models.Color, error)
}

type ColorRepo struct {
	db *database.DB
}

const (
	ColorTable = "colors"
)

var _ ColorRepository = (*ColorRepo)(nil)

func NewColorRepo(db *database.DB) ColorRepository {
	return &ColorRepo{db: db}
}

func (repo *ColorRepo) Create(m *models.CreateColor) (string, error) {
	query := fmt.Sprintf(`INSERT INTO "%s" (color,name) VALUES ($1,$2) RETURNING id`, ColorTable)
	newColor := repo.db.QueryRow(query, m.Color, m.Name)
	var colorId string
	if err := newColor.Scan(&colorId); err != nil {
		return "", err
	}

	return colorId, nil
}

func (repo *ColorRepo) GetColorById(colorId uuid.UUID) (*models.Color, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, ColorTable)
	color := &models.Color{}
	err := repo.db.Get(color, query, colorId)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return color, nil
}

func (repo *ColorRepo) All() ([]*models.Color, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" `, ColorTable)
	colorList := &[]*models.Color{}
	err := repo.db.Select(colorList, query)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return *colorList, nil
}

func (repo *ColorRepo) UpdateColor(m *models.UpdateColor) error {
	query := fmt.Sprintf(`UPDATE "%s" SET name=$2, color=$3 WHERE id =$1`, ColorTable)
	if _, err := repo.db.Exec(query, m.ID, m.Name, m.Color); err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNotFound
		}
		return err
	}

	return nil
}
