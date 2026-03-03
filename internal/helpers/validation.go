package helpers

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var validationErrors validator.ValidationErrors
	var errorMessages []string

	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			switch e.Tag() {

			case "required":
				errorMessages = append(errorMessages,
					fmt.Sprintf("%s harus diisi", e.Field()))

			case "min":
				errorMessages = append(errorMessages,
					fmt.Sprintf("%s minimal %s karakter", e.Field(), e.Param()))

			case "max":
				errorMessages = append(errorMessages,
					fmt.Sprintf("%s maksimal %s karakter", e.Field(), e.Param()))

			default:
				errorMessages = append(errorMessages,
					fmt.Sprintf("%s tidak valid", e.Field()))
			}
		}
	}

	return errorMessages
}
