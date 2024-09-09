package page

import "context"

type PaginationHelper[T, S any] interface {
	Pagination(ctx context.Context, p Paginator[T], filter S) error
}

// Paginator defines the interface for pagination operations.
type Paginator[T any] interface {
	GetCurrent() int
	SetCurrent(current int)
	GetPageSize() int
	SetPageSize(pageSize int)
	GetTotal() int64
	SetTotal(total int64)
	GetItems() []T
	SetItems(items []T)
	Offset() int
}

// pagination is a generic pagination structure to hold paginated data of any type.
type pagination[T any] struct {
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
	Items    []T   `json:"items,omitempty"`
}

// NewPagination creates a new pagination instance with the provided current page and page size.
// It ensures the page number and page size are always valid.
func NewPagination[T any](current, pageSize int) *pagination[T] {
	// Initialize current and pageSize with default values if invalid values are provided.
	if current < 1 {
		current = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return &pagination[T]{
		Current:  current,
		PageSize: pageSize,
	}
}

// GetCurrent returns the current page number.
func (p *pagination[T]) GetCurrent() int {
	return p.Current
}

// SetCurrent safely sets the current page number.
func (p *pagination[T]) SetCurrent(current int) {
	if current > 0 {
		p.Current = current
	}
}

// GetPageSize returns the number of items per page.
func (p *pagination[T]) GetPageSize() int {
	return p.PageSize
}

// SetPageSize safely sets the number of items per page.
func (p *pagination[T]) SetPageSize(pageSize int) {
	if pageSize > 0 {
		p.PageSize = pageSize
	}
}

// GetTotal returns the total number of items.
func (p *pagination[T]) GetTotal() int64 {
	return p.Total
}

// SetTotal safely sets the total number of items.
func (p *pagination[T]) SetTotal(total int64) {
	p.Total = total
}

// GetItems returns the items in the current page.
func (p *pagination[T]) GetItems() []T {
	return p.Items
}

// SetItems sets the items in the current page.
func (p *pagination[T]) SetItems(items []T) {
	p.Items = items
}

// Offset calculates the start offset for the current page based on the current page number and page size.
func (p *pagination[T]) Offset() int {
	return (p.Current - 1) * p.PageSize
}
