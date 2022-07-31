package shared

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type FilterConditions interface {
	WithFields() bool
	OnlyFields(prefix string) []string
	HasConditions() bool
	Conditions() (string, []interface{})
	GetCondition(key string) (interface{}, bool)
}

type Filters struct {
	Fields       []string
	ConditionMap map[string]interface{}
}

func (f *Filters) OnlyFields(prefix string) []string {
	if prefix != "" {
		withPrefix := make([]string, len(f.Fields))
		for i, field := range f.Fields {
			withPrefix[i] = fmt.Sprintf("%s.%s", prefix, field)
		}

		return withPrefix
	}

	return f.Fields
}

func (f *Filters) WithFields() bool {
	return len(f.Fields) > 0
}

func (f *Filters) HasConditions() bool {
	return len(f.ConditionMap) > 0
}

func (f *Filters) Conditions() (string, []interface{}) {
	statement := []string{}
	values := []interface{}{}

	// f.conditions looks like:
	// map[
	// 	"course_published": true,
	// 	"course_name":      "Effective Eureka",
	// ]
	//
	// So key is course_name as example and value is "Effective Eureka"
	for key, value := range f.ConditionMap {
		if key != "" {
			// statement looks like:
			// []string{"course_published = ?", "course_name = ?"}
			statement = append(statement, key+" = ?")
			// values looks like:
			// []interface{}{true, "Effective Eureka"}
			values = append(values, value)
		}
	}

	// finalStatement looks like:
	// "course_published = ? AND course_name = ?"
	finalStatement := strings.Join(statement, " AND ")

	return finalStatement, values
}

func (f *Filters) GetCondition(key string) (interface{}, bool) {
	if cond, ok := f.ConditionMap[key]; ok {
		return cond, ok
	}

	return nil, false
}

type Paginator interface {
	ShouldPaginate() bool
	GetOffset() uint16
	GetItemsPerPage() int
	Paginate(db *gorm.DB) *gorm.DB
}

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
