package repository

import (
	"fmt"
	"strings"
)

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
