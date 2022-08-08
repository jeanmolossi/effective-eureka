package paginator

import (
	"gorm.io/gorm"
)

type paginator struct {
	baseURL      string
	page         uint16
	itemsPerPage int
}

func NewPaginator() Paginator {
	return &paginator{"", 1, 10}
}

func (p *paginator) SetBaseURL(baseURL string) {
	p.baseURL = baseURL
}

func (p *paginator) SetPage(page uint16) {
	if page == 0 {
		page = 1
	}

	if page >= 65535 {
		page = 65535
	}

	p.page = page
}

func (p *paginator) SetItemsPerPage(itemsPerPage int) {
	if itemsPerPage == 0 || itemsPerPage > 100 {
		itemsPerPage = 10
	}

	p.itemsPerPage = itemsPerPage
}

func (p *paginator) Page() uint16 {
	return p.page
}

func (p *paginator) GetOffset() uint16 {
	return uint16(p.itemsPerPage) * (p.page - 1)
}

func (p *paginator) GetItemsPerPage() int {
	return p.itemsPerPage
}

func (p *paginator) ShouldPaginate() bool {
	return p.page > 1 || p.itemsPerPage != 10
}

func (p *paginator) Paginate(db *gorm.DB) *gorm.DB {
	offset := int(p.GetOffset())
	return db.Offset(offset).Limit(p.GetItemsPerPage())
}
