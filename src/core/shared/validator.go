package shared

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ModuleErrorMap is an interface with input structs of data should implements
// to get validated with custom validator.
type ModuleErrorMap interface {
	// GetErrorMap is the method who returns a map with errors.
	GetErrorMap() map[string]map[string]error
}

// CustomValidator is a struct to manage the validator.
type CustomValidator struct {
	validate *validator.Validate
}

// ValidationErr is a struct to manage the validation errors.
// Err is the main error message.
// Errors is a slice of errors. That slice will contain the errors of every
// field validation.
type ValidationErr struct {
	Err    string  `json:"error"`
	Errors []error `json:"errors,omitempty"`
}

// AddError is a method to add an error to the slice of errors.
func (v *ValidationErr) AddError(err error) {
	v.Errors = append(v.Errors, err)
}

// Error implements error interface.
func (v *ValidationErr) Error() string {
	return fmt.Sprintf("404: %s", v.Err)
}

// NewCustomValidation is a factory method to create a CustomValidator.
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validate: validator.New(),
	}
}

// Validate is a method to validate a struct with custom validator.
// Param `i` should be an struct tagged with validate tag.
func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validate.Struct(i)

	// Default bad request validation error
	fieldErrors := &ValidationErr{
		Err:    "Bad Request",
		Errors: []error{},
	}

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println("Validation error:", err)
			return err
		}

		// cast our param as ModuleErrorMap
		if errMap, ok := i.(ModuleErrorMap); ok {
			// get the map of errors
			for _, err := range err.(validator.ValidationErrors) {
				if fieldErr := GetFieldError(err, errMap); fieldErr != nil {
					// when we have a field error, add it to the slice of errors
					fieldErrors.AddError(fieldErr)
				}
			}
		} else {
			// If we can't cast the param as ModuleErrorMap, we return the default error
			return errors.New("not module error map")
		}

		return fieldErrors
	}

	return nil
}

// GetFieldError gets a validation field error and ModuleErrorMap.
// That function select the correct field error from the map of errors and
// return that error.
func GetFieldError(err validator.FieldError, moduleErrMap ModuleErrorMap) error {
	errMap := moduleErrMap.GetErrorMap()

	// set the field as lowercase.
	// Example: Title will be title
	lowField := strings.ToLower(err.Field())
	// Gets error tag. If the error is because required tag, or max tag...
	// Example: required
	errTag := err.Tag()

	// If the field error exists in the map, we return the error
	if errMap[lowField] != nil {
		return &FieldError{
			Field: lowField,
			Err:   errMap[lowField][errTag].Error(),
		}
	}

	// If the field error doesn't exist in the map, we return the default error
	return &FieldError{Field: lowField, Err: "unknown error"}
}

// FieldError is a struct to manage the field errors.
type FieldError struct {
	Field string `json:"field" example:"title"`
	Err   string `json:"message" example:"title is required"`
}

// Error implements error interface.
func (f *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", f.Field, f.Err)
}
