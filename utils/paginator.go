package utils

type PaginatorParams struct {
	Page     int
	PageSize int
}

func GetPaginatorParams(defaultPage int, defaultPageSize int, a []interface{}) (*PaginatorParams, error) {
	page := defaultPage
	pageSize := defaultPageSize

	if len(a) > 0 {
		page = GetAtoi(a[0])
	}

	if len(a) > 1 {
		pageSize = GetAtoi(a[1])
	}

	if page < 1 {
		page = defaultPage
	}

	if pageSize < 1 {
		pageSize = defaultPageSize
	}

	return &PaginatorParams{Page: page, PageSize: pageSize}, nil
}

func (params *PaginatorParams) GetOffset() int {
	return (params.Page - 1) * params.PageSize
}
