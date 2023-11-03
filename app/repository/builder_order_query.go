package repository

import (
	"fmt"
	"strings"
)

type orderQueryBuilder struct {
	query  string
	sortBy string
	sort   string
}

func (*OrderRepo) newOrderQueryBuilder(query string) *orderQueryBuilder {
	return &orderQueryBuilder{
		query:  query,
		sortBy: "id",
		sort:   "ASC",
	}
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
	q.query += fmt.Sprintf("ORDER BY %s %s", q.sortBy, q.sort)
	return q.query
}
