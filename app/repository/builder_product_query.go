package repository

import (
	"fmt"

	"github.com/google/uuid"
)

type ProductQueryBuilder struct {
	query  string
	field  map[string]bool
	sortBy string
	sort   string
}

func NewProductQueryBuilder(query string) *ProductQueryBuilder {
	return &ProductQueryBuilder{
		query:  query,
		field:  make(map[string]bool),
		sortBy: "id",
		sort:   "ASC",
	}
}

func (q *ProductQueryBuilder) SetName(name string) *ProductQueryBuilder {
	if name != "" {
		query := fmt.Sprintf(` word_similarity(published_products."name",'%s') > 0 `, name)
		q.field[query] = true
	}

	return q
}

func (q *ProductQueryBuilder) SetAvailable(available *bool) *ProductQueryBuilder {
	if available != nil {
		query := fmt.Sprintf(` available = '%v'`, available)
		q.field[query] = true
	}

	return q
}

func (q *ProductQueryBuilder) SetCategory(category *uuid.UUID) *ProductQueryBuilder {
	if category != nil {
		query := fmt.Sprintf(` category_id = '%s'`, category)
		q.field[query] = true
	}

	return q
}

func (q *ProductQueryBuilder) SetSize(size string) *ProductQueryBuilder {
	if size != "" {
		query := fmt.Sprintf(`size='%s'`, size)
		q.field[query] = true
	}

	return q
}

func (q *ProductQueryBuilder) SetType(types string) *ProductQueryBuilder {
	if types != "" {
		query := fmt.Sprintf("type= '%s'", types)
		q.field[query] = true
	}

	return q
}

func (q *ProductQueryBuilder) SetDifficulty(difficulty string) *ProductQueryBuilder {
	if difficulty != "" {
		query := fmt.Sprintf("difficulty = '%s'", difficulty)
		q.field[query] = true
	}

	return q
}

func (q *ProductQueryBuilder) SetWarter(warter string) *ProductQueryBuilder {
	if warter != "" {
		query := fmt.Sprintf("warter = '%s'", warter)
		q.field[query] = true
	}

	return q
}

func (q *ProductQueryBuilder) SortBy(field string, sort string) *ProductQueryBuilder {
	if field != "" {
		q.sortBy = field
	}
	if sort == "ASC" || sort == "DESC" {
		q.sort = sort
	}

	return q
}

func (q *ProductQueryBuilder) Build() string {
	count := 0
	if len(q.field) != 0 {
		q.query += "WHERE "
		for k := range q.field {
			if count != 0 {
				q.query += "AND "
			}
			q.query += k

			count++
		}
	}

	q.query += fmt.Sprintf("ORDER BY %s %s", q.sortBy, q.sort)
	return q.query
}
