package repository

import (
	"fmt"
	"strings"
)

type cartProductQueryBuilder struct {
	query  string
	sortBy string
	sort   string
}

func (*CartRepo) newCartQueryBuilder(query string) *cartProductQueryBuilder {
	return &cartProductQueryBuilder{
		query:  query,
		sortBy: "id",
		sort:   "ASC",
	}
}

func (q *cartProductQueryBuilder) SortBy(field string, sort string) *cartProductQueryBuilder {
	if field != "" {
		q.sortBy = field
	}
	sort = strings.ToUpper(sort)
	if sort == "ASC" || sort == "DESC" {
		q.sort = sort
	}

	return q
}

func (q *cartProductQueryBuilder) Build() string {
	q.query += fmt.Sprintf("ORDER BY %s %s", q.sortBy, q.sort)
	return q.query
}
