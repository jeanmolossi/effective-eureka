package paginator

import "gorm.io/gorm"

type Paginator interface {
	Page() uint16

	SetBaseURL(baseURL string)
	SetPage(page uint16)
	SetItemsPerPage(itemsPerPage int)

	ShouldPaginate() bool
	GetOffset() uint16
	GetItemsPerPage() int

	Paginate(db *gorm.DB) *gorm.DB
}
