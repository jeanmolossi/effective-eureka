package shared

type Event string

const (
	CourseCreated      Event = "course_created"
	ModuleCreated      Event = "module_created"
	LessonCreated      Event = "lesson_created"
	LessonItemAttached Event = "lesson_item_attached"
	LessonItemDetached Event = "lesson_item_detached"
)
