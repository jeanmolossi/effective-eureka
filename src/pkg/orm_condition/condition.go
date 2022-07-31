package ormcondition

type FilterConditions interface {
	WithFields() bool
	SelectFields(prefix string) []string
	HasConditions() bool
	Conditions() (string, []interface{})
	GetCondition(key string) (interface{}, bool)

	AddCondition(field string, value interface{})
	RemoveCondition(field string)
	AddField(field string)
	AddFields(fields []string)
}
