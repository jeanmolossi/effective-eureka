package shared

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ModuleErrorMap interface {
	GetErrorMap() map[string]map[string]error
}

type CustomValidator struct {
	validate *validator.Validate
}

type ValidationErr struct {
	Err    string  `json:"error"`
	Errors []error `json:"errors,omitempty"`
}

func (v *ValidationErr) AddError(err error) {
	v.Errors = append(v.Errors, err)
}

func (v *ValidationErr) Error() string {
	return fmt.Sprintf("404: %s", v.Err)
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validate: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validate.Struct(i)

	fieldErrors := &ValidationErr{
		Err:    "Bad Request",
		Errors: []error{},
	}

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println("Validation error:", err)
			return err
		}

		if errMap, ok := i.(ModuleErrorMap); ok {
			for _, err := range err.(validator.ValidationErrors) {
				if fieldErr := GetFieldError(err, errMap); fieldErr != nil {
					fieldErrors.AddError(fieldErr)
				}
			}
		} else {
			return errors.New("not module error map")
		}

		return fieldErrors
	}

	return nil
}

func GetFieldError(err validator.FieldError, moduleErrMap ModuleErrorMap) error {
	errMap := moduleErrMap.GetErrorMap()

	lowField := strings.ToLower(err.Field())
	errTag := err.Tag()

	fmt.Println(lowField, errTag)

	if errMap[lowField] != nil {
		return &FieldError{
			Field: lowField,
			Err:   errMap[lowField][errTag].Error(),
		}
	}

	return &FieldError{Field: lowField, Err: "unknown error"}
}

type FieldError struct {
	Field string `json:"field"`
	Err   string `json:"message"`
}

func (f *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", f.Field, f.Err)
}
