package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)


type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
func BindAndValidate(c echo.Context, req interface{}) error {

	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	return nil
}
