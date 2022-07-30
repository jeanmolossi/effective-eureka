package repository

import "gorm.io/gorm"

type PagesConfig struct {
	Page         uint16
	ItemsPerPage int
}

func (p *PagesConfig) GetOffset() uint16 {
	return uint16(p.ItemsPerPage) * (p.Page - 1)
}

func (p *PagesConfig) GetItemsPerPage() int {
	return p.ItemsPerPage
}

func (p *PagesConfig) ShouldPaginate() bool {
	return true
}

func (p *PagesConfig) Paginate(db *gorm.DB) *gorm.DB {
	offset := int(p.GetOffset())
	return db.Offset(offset).Limit(p.GetItemsPerPage())
}
