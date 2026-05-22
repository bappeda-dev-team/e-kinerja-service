package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) string {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return err.Error()
	}

	msgs := make([]string, 0, len(validationErrors))
	for _, e := range validationErrors {
		switch e.Tag() {
		case "required":
			msgs = append(msgs, fmt.Sprintf("%s harus diisi", e.Field()))
		case "min":
			msgs = append(msgs, fmt.Sprintf("%s minimal %s karakter", e.Field(), e.Param()))
		case "max":
			msgs = append(msgs, fmt.Sprintf("%s maksimal %s karakter", e.Field(), e.Param()))
		default:
			msgs = append(msgs, fmt.Sprintf("%s tidak valid", e.Field()))
		}
	}
	return strings.Join(msgs, "; ")
}
