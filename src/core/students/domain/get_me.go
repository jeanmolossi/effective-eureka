package domain

type GetMeParams struct {
	StudentID string   `json:"student_id"`
	Fields    []string `query:"fields"`
}

// GetMe returns the auth user.
type GetMe interface {
	// Run execute get me.
	Run(input *GetMeParams) (Student, error)
}
