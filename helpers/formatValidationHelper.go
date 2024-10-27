// helpers/formatHelper.go
package helpers

import (
    "github.com/go-playground/validator/v10"
)

// FormatValidationError takes a validation error and returns a map of field names with their respective error messages.
func FormatValidationError(err error) map[string]string {
    errors := map[string]string{}

    // Check if the error is related to validation
    if _, ok := err.(*validator.InvalidValidationError); ok {
        return errors
    }

    // Loop over validation errors and format them
    for _, err := range err.(validator.ValidationErrors) {
        errors[err.Field()] = "failed validation with rule: " + err.Tag()
    }
    return errors
}
