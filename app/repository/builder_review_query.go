package repository

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type reviewQueryBuilder struct {
	query  string
	field  map[string]bool
	sortBy string
	sort   string
}

func (*ReviewRepo) newReviewQueryBuilder(query string) *reviewQueryBuilder {
	return &reviewQueryBuilder{
		query:  query,
		field:  make(map[string]bool),
		sortBy: "created_at",
		sort:   " ASC",
	}
}

func (q *reviewQueryBuilder) SetProduct(productId *uuid.UUID) *reviewQueryBuilder {
	if productId != nil {
		query := fmt.Sprintf(` reviews.product_id = '%s'`, productId)
		q.field[query] = true
	}

	return q
}

func (q *reviewQueryBuilder) SetStar(star int) *reviewQueryBuilder {
	if star > 0 {
		query := fmt.Sprintf(` reviews.star = %d`, star)
		q.field[query] = true
	}

	return q
}

func (q *reviewQueryBuilder) SetUser(userId *uuid.UUID) *reviewQueryBuilder {
	if userId != nil {
		query := fmt.Sprintf(` reviews.user_id = '%s'`, userId)
		q.field[query] = true
	}

	return q
}

func (q *reviewQueryBuilder) SortBy(field string, sort string) *reviewQueryBuilder {
	if field != "" {
		q.sortBy = field
	}

	sort = strings.ToUpper(sort)
	if sort == "ASC" || sort == "DESC" {
		q.sort = " " + sort + " "
	}

	return q
}

func (q *reviewQueryBuilder) Build() string {
	count := 0
	if len(q.field) != 0 {
		q.query += "WHERE "
		for k := range q.field {
			if count != 0 {
				q.query += " AND "
			}
			q.query += k

			count++
		}
	}

	q.query += fmt.Sprintf(" ORDER BY %s %s", q.sortBy, q.sort)
	return q.query
}
