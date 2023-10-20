package models

type BaseQuery struct {
	OffSet int    `query:"offSet"`
	Limit  int    `query:"limit"`
	Sort   string `query:"sort"`
	SortBy string `query:"sortBy"`
}

func DefaultQuery() *BaseQuery {
	return &BaseQuery{
		OffSet: 1,
		Limit:  10,
	}
}

type BasePaginationResponse struct {
	Items    any  `json:"items"`
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	Next     bool `json:"next"`
	Prev     bool `json:"prev"`
}

func (q *BaseQuery) IsFirstPage() bool {
	return q.OffSet < 2
}

func (q *BaseQuery) HaveNextPage(arrLen int) bool {
	if q.Limit <= 0 {
		return false
	}

	return arrLen > q.Limit
}

func (q *BaseQuery) GetPageNumber() int {
	return (q.OffSet + q.Limit - 1) / q.Limit
}
