package models

type Pagination struct {
	OffSet int
	Limit  int
}

func DefaultPagination() *Pagination {
	return &Pagination{
		OffSet: 1,
		Limit:  10,
	}
}
