package repository

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type orderQueryBuilder struct {
	query  string
	sortBy string
	field  map[string]bool
	sort   string
}

func (*OrderRepo) newOrderQueryBuilder(query string) *orderQueryBuilder {
	return &orderQueryBuilder{
		query:  query,
		field:  make(map[string]bool),
		sortBy: "id",
		sort:   "ASC",
	}
}

func (q *orderQueryBuilder) SetState(state string) *orderQueryBuilder {
	if state != "" {
		query := fmt.Sprintf(`state = '%s'`, state)
		q.field[query] = true
	}

	return q
}

func (q *orderQueryBuilder) SetCoupon(couponId *uuid.UUID) *orderQueryBuilder {
	if couponId != nil {
		query := fmt.Sprintf(`coupon_id = '%s'`, couponId)
		q.field[query] = true
	}

	return q
}

func (q *orderQueryBuilder) SetOwner(ownerId *uuid.UUID) *orderQueryBuilder {
	if ownerId != nil {
		query := fmt.Sprintf(`owner_id = '%s'`, ownerId)
		q.field[query] = true
	}
	return q
}

func (q *orderQueryBuilder) SortBy(field string, sort string) *orderQueryBuilder {
	if field != "" {
		q.sortBy = field
	}
	sort = strings.ToUpper(sort)
	if sort == "ASC" || sort == "DESC" {
		q.sort = sort
	}

	return q
}

func (q *orderQueryBuilder) Build() string {
	count := 0
	if len(q.field) != 0 {
		q.query += " WHERE "
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
