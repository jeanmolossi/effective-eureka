package shared

type FilterConditions interface {
	WithFields() bool
	OnlyFields(prefix string) []string
	HasConditions() bool
	Conditions() (string, []interface{})
}
