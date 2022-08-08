package domain

import "errors"

type GetLessonParams struct {
	LessonID string   `param:"lessonID" validate:"required"`
	Fields   []string `query:"fields"`
}

func (c *GetLessonParams) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"lessonid": {
			"required": errors.New("lesson id is required"),
		},
	}
}

type GetLesson interface {
	Run(params *GetLessonParams) (Lesson, error)
}
