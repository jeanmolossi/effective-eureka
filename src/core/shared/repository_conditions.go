package shared

import "gorm.io/gorm"

type FilterConditions interface {
	WithFields() bool
	OnlyFields(prefix string) []string
	HasConditions() bool
	Conditions() (string, []interface{})
	GetCondition(key string) (interface{}, bool)
}

type Paginator interface {
	ShouldPaginate() bool
	GetOffset() uint16
	GetItemsPerPage() int
	Paginate(db *gorm.DB) *gorm.DB
}
