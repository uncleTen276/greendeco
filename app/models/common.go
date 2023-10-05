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
	Items    any `json:"items"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}
