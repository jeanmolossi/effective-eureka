package ormcondition

import (
	"fmt"
	"strings"
)

type filters struct {
	fields       []string
	conditionMap map[string]interface{}
}

func NewFilterConditions() FilterConditions {
	return &filters{conditionMap: make(map[string]interface{})}
}

func (f *filters) SelectFields(prefix string) []string {
	if prefix != "" {
		withPrefix := make([]string, len(f.fields))
		for i, field := range f.fields {
			withPrefix[i] = fmt.Sprintf("%s.%s", prefix, field)
		}

		return withPrefix
	}

	return f.fields
}

func (f *filters) WithFields() bool {
	return len(f.fields) > 0
}

func (f *filters) HasConditions() bool {
	return len(f.conditionMap) > 0
}

func (f *filters) Conditions() (string, []interface{}) {
	statement := []string{}
	values := []interface{}{}

	// f.conditions looks like:
	// map[
	// 	"course_published": true,
	// 	"course_name":      "Effective Eureka",
	// ]
	//
	// So key is course_name as example and value is "Effective Eureka"
	for key, value := range f.conditionMap {
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

func (f *filters) GetCondition(key string) (interface{}, bool) {
	if cond, ok := f.conditionMap[key]; ok {
		return cond, ok
	}

	return nil, false
}

func (f *filters) AddCondition(field string, value interface{}) {
	f.conditionMap[field] = value
}

func (f *filters) RemoveCondition(field string) {
	delete(f.conditionMap, field)
}

func (f *filters) AddField(field string) {
	f.fields = append(f.fields, field)
}

func (f *filters) AddFields(fields []string) {
	if len(fields) == 0 {
		return
	}

	for _, field := range fields {
		f.AddField(field)
	}
}
