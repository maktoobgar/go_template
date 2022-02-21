package errors

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	serverError struct {
		code    int
		action  int
		message string
	}
)

const (
	_ int = iota
	// BadRequest 400
	InvalidStatus
	// NotFound 404
	NotFoundStatus
	// Unauthorized 401
	UnauthorizedStatus
	// InternalServerError 500
	UnexpectedStatus
	// MethodNotAllowed 405
	NotAllowedStatus
)

const (
	_ int = iota
	// SignIn in again
	ReSingIn
	// Report the problem
	Report
	// Correct sent data and request again
	Resend
)

var (
	httpErrors = map[int]int{
		InvalidStatus:      http.StatusBadRequest,
		NotFoundStatus:     http.StatusNotFound,
		UnauthorizedStatus: http.StatusUnauthorized,
		UnexpectedStatus:   http.StatusInternalServerError,
		NotAllowedStatus:   http.StatusMethodNotAllowed,
	}
)

func (e serverError) Error() string {
	return e.message
}

// Returns httpErrorCode, message and action of it
func HttpError(err error) (code int, action int, message string) {
	code = http.StatusInternalServerError
	action = Report
	message = err.Error()

	if er, ok := err.(serverError); ok {
		code = httpErrors[er.code]
		action = er.action
		message = er.message
	}

	return
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Set Content-Type: application/json; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	code, action, message := HttpError(err)

	return c.Status(code).JSON(map[string]string{
		"Code":    fmt.Sprint(code),
		"Action":  fmt.Sprint(action),
		"Message": message,
	})
}

// Creates a new error
func New(code int, action int, message string) error {
	return serverError{
		code:    code,
		action:  action,
		message: message,
	}
}
