package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Errors  map[string]string `json:"errors"`
	Message string            `json:"message,omitempty"`
}

var defaultErrorMessages = map[string]string{
	"required": "is required",
	"min":      "requires a minimum length of",
	"max":      "requires a maximum length of",
	"email":    "is not a valid email address",
	"url":      "is not a valid URL",
	"numeric":  "must be a number",
}

func NewErrorResponse(err error) ErrorResponse {
	errors := make(map[string]string)
	message := "Validation failed for one or more fields."

	// The dot operator is used for type assertion, ok is true if the dynamic type
	// of err cam be assigned to the given error, in this case validator.ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			fieldName := strings.ToLower(fieldErr.Field())

			msg, found := defaultErrorMessages[fieldErr.Tag()]
			if !found {
				msg = fmt.Sprintf("has an invalid value (rule: %s)", fieldErr.Tag())
			}

			if fieldErr.Tag() == "min" || fieldErr.Tag() == "max" || fieldErr.Tag() == "len" {
				msg = fmt.Sprintf("%s %s", msg, fieldErr.Param())
			}

			errors[fieldName] = fmt.Sprintf("%s %s", fieldName, msg)
		}
	} else {
		errors["unknown"] = "An unexpected validation error occurred."
		message = "An unexpected error occurred during validation."
	}

	return ErrorResponse{
		Errors:  errors,
		Message: message,
	}
}
