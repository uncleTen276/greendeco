package repository

import (
	"fmt"
	"strings"
)

type queryBuilder struct {
	query  string
	sortBy string
	sort   string
}

func newQueryBuilder(query string, baseField string) *queryBuilder {
	return &queryBuilder{
		query:  query,
		sortBy: baseField,
		sort:   "ASC",
	}
}

func (q *queryBuilder) SortBy(field string, sort string) *queryBuilder {
	if field != "" {
		q.sortBy = field
	}

	sort = strings.ToUpper(sort)
	if sort == "ASC" || sort == "DESC" {
		q.sort = sort
	}

	return q
}

func (q *queryBuilder) Build() string {
	q.query += fmt.Sprintf(" ORDER BY %s %s", q.sortBy, q.sort)
	return q.query
}
