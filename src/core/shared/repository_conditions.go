package shared

import "gorm.io/gorm"

type FilterConditions interface {
	WithFields() bool
	OnlyFields(prefix string) []string
	HasConditions() bool
	Conditions() (string, []interface{})
}

type Paginator interface {
	ShouldPaginate() bool
	GetOffset() uint16
	GetItemsPerPage() int
	Paginate(db *gorm.DB) *gorm.DB
}
