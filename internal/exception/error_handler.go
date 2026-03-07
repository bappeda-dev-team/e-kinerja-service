package exception

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ============================================================
// HTTP EXCEPTION INTERFACE & BASE
// ============================================================

type HttpException interface {
	error
	StatusCode() int
	GetMessage() string
}

type baseException struct {
	Code    int
	Message string
}

func (e *baseException) Error() string      { return e.Message }
func (e *baseException) StatusCode() int    { return e.Code }
func (e *baseException) GetMessage() string { return e.Message }

// ============================================================
// CONSTRUCTORS
// ============================================================

// 400
func BadRequest(msg string) *baseException   { return &baseException{http.StatusBadRequest, msg} }
func InvalidInput(msg string) *baseException { return &baseException{http.StatusBadRequest, msg} }
func Validation(msg string) *baseException   { return &baseException{http.StatusBadRequest, msg} }

// 401
func Unauthentication(msg string) *baseException { return &baseException{http.StatusUnauthorized, msg} }
func InvalidToken(msg string) *baseException     { return &baseException{http.StatusUnauthorized, msg} }
func Unauthorized(msg string) *baseException     { return &baseException{http.StatusUnauthorized, msg} }

// 403
func Forbidden(msg string) *baseException    { return &baseException{http.StatusForbidden, msg} }
func AccessDenied(msg string) *baseException { return &baseException{http.StatusForbidden, msg} }

// 404
func ResourceNotFound(msg string) *baseException { return &baseException{http.StatusNotFound, msg} }

// 409
func Conflict(msg string) *baseException          { return &baseException{http.StatusConflict, msg} }
func DuplicateResource(msg string) *baseException { return &baseException{http.StatusConflict, msg} }

// 429
func RateLimit(msg string) *baseException { return &baseException{http.StatusTooManyRequests, msg} }

// 500
func InternalServer(msg string) *baseException {
	return &baseException{http.StatusInternalServerError, msg}
}
func Database(msg string) *baseException {
	return &baseException{http.StatusInternalServerError, "Database operation failed. Please try again later."}
}
func Service(msg string) *baseException { return &baseException{http.StatusInternalServerError, msg} }

// ============================================================
// API RESPONSE STRUCTURE
// ============================================================

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func errorResponse(status int, message string, data interface{}) ApiResponse {
	return ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// ============================================================
// GLOBAL ERROR HANDLER (Echo)
// ============================================================

// GlobalExceptionHandler returns an echo.HTTPErrorHandler
// Usage: e.HTTPErrorHandler = exception.GlobalExceptionHandler()
func GlobalExceptionHandler() func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		// Handle Echo's own HTTPError (404 route not found, 405 method not allowed, etc.)
		var he *echo.HTTPError
		if errors.As(err, &he) {
			msg, ok := he.Message.(string)
			if !ok {
				msg = "Unknown error"
			}
			_ = c.JSON(he.Code, errorResponse(he.Code, msg, nil))
			return
		}

		// Handle custom exceptions
		var httpErr HttpException
		if errors.As(err, &httpErr) {
			_ = c.JSON(httpErr.StatusCode(), errorResponse(httpErr.StatusCode(), httpErr.GetMessage(), nil))
			return
		}

		// Fallback
		_ = c.JSON(http.StatusInternalServerError, errorResponse(
			http.StatusInternalServerError,
			"An unexpected error occurred. Please contact support.",
			nil,
		))
	}
}
