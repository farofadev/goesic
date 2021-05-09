package utils

import "fmt"

type PaginatorParams struct {
	Page     int
	PageSize int
}

type PaginatorConfig struct {
	MaxPage         int
	MaxPageSize     int
	DefaultPage     int
	DefaultPageSize int
}

const (
	DefaultPaginatorPageSize    = 25
	DefaultPagaintorMaxPageSize = 100
	DefaultPaginatorPage        = 1
)

func NewPaginatorConfig() *PaginatorConfig {
	return &PaginatorConfig{
		MaxPage:         -1,
		MaxPageSize:     DefaultPagaintorMaxPageSize,
		DefaultPage:     DefaultPaginatorPage,
		DefaultPageSize: DefaultPaginatorPageSize,
	}
}

func GetPaginatorParams(config *PaginatorConfig, a ...interface{}) (*PaginatorParams, error) {
	page := config.DefaultPage
	pageSize := config.DefaultPageSize

	if len(a) > 0 {
		page = GetAtoi(a[0])
	}

	if len(a) > 1 {
		pageSize = GetAtoi(a[1])
	}

	if page < 1 {
		page = config.DefaultPage
	}

	if pageSize < 1 {
		pageSize = config.DefaultPageSize
	}

	if config.MaxPage > -1 && page > config.MaxPage {
		page = config.MaxPage
	}

	if config.MaxPageSize > -1 && pageSize > config.MaxPageSize {
		pageSize = config.MaxPageSize
	}

	return &PaginatorParams{Page: page, PageSize: pageSize}, nil
}

func (params *PaginatorParams) GetOffset() int {
	return (params.Page - 1) * params.PageSize
}

func (params *PaginatorParams) GetCacheKey() string {
	return fmt.Sprintf("page_%d_size_%d", params.Page, params.PageSize)
}
