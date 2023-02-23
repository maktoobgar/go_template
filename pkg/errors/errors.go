package errors

import (
	"net/http"
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
	MethodNotAllowedStatus
	// Forbidden 403
	ForbiddenStatus
)

const (
	// Do nothing
	DoNothing int = iota
	// SignIn in again
	ReSignIn
	// Report the problem
	Report
	// Correct sent data and request again
	Resend
)

var (
	httpErrors = map[int]int{
		InvalidStatus:          http.StatusBadRequest,
		NotFoundStatus:         http.StatusNotFound,
		UnauthorizedStatus:     http.StatusUnauthorized,
		UnexpectedStatus:       http.StatusInternalServerError,
		MethodNotAllowedStatus: http.StatusMethodNotAllowed,
		ForbiddenStatus:        http.StatusForbidden,
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

func IsServerError(err error) bool {
	if _, ok := err.(serverError); ok {
		return true
	}
	return false
}

// Creates a new error
func New(code int, action int, message string) error {
	return serverError{
		code:    code,
		action:  action,
		message: message,
	}
}
