package repository

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type productQueryBuilder struct {
	query  string
	field  map[string]bool
	sortBy string
	sort   string
}

func (*ProductRepo) newProductQueryBuilder(query string) *productQueryBuilder {
	return &productQueryBuilder{
		query:  query,
		field:  make(map[string]bool),
		sortBy: "id",
		sort:   "ASC",
	}
}

func (q *productQueryBuilder) SetName(name string) *productQueryBuilder {
	if name != "" {
		query := fmt.Sprintf(` word_similarity(published_products."name",'%s') > 0 `, name)
		q.field[query] = true
	}

	return q
}

func (q *productQueryBuilder) SetAvailable(available *bool) *productQueryBuilder {
	if available != nil {
		query := fmt.Sprintf(` available = '%v'`, available)
		q.field[query] = true
	}

	return q
}

func (q *productQueryBuilder) SetCategory(category *uuid.UUID) *productQueryBuilder {
	if category != nil {
		query := fmt.Sprintf(` category_id = '%s'`, category)
		q.field[query] = true
	}

	return q
}

func (q *productQueryBuilder) SetSize(size string) *productQueryBuilder {
	if size != "" {
		query := fmt.Sprintf(`size='%s'`, size)
		q.field[query] = true
	}

	return q
}

func (q *productQueryBuilder) SetType(types string) *productQueryBuilder {
	if types != "" {
		query := fmt.Sprintf("type= '%s'", types)
		q.field[query] = true
	}

	return q
}

func (q *productQueryBuilder) SetDifficulty(difficulty string) *productQueryBuilder {
	if difficulty != "" {
		query := fmt.Sprintf("difficulty = '%s'", difficulty)
		q.field[query] = true
	}

	return q
}

func (q *productQueryBuilder) Setwater(water string) *productQueryBuilder {
	if water != "" {
		query := fmt.Sprintf("water = '%s'", water)
		q.field[query] = true
	}

	return q
}

func (q *productQueryBuilder) SortBy(field string, sort string) *productQueryBuilder {
	if field != "" {
		q.sortBy = field
	}
	sort = strings.ToUpper(sort)
	if sort == "ASC" || sort == "DESC" {
		q.sort = sort
	}

	return q
}

func (q *productQueryBuilder) Build() string {
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
