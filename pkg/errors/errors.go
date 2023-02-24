package errors

import (
	"net/http"
)

type (
	serverError struct {
		code    int
		action  int
		message string
		errors  any
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
	// Timeout 408
	TimeoutStatus
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
		TimeoutStatus:          http.StatusRequestTimeout,
	}
)

func (e serverError) Error() string {
	return e.message
}

// Returns httpErrorCode, message and action of it
func HttpError(err error) (code int, action int, message string, errors any) {
	code = http.StatusInternalServerError
	action = Report
	message = err.Error()
	errors = nil

	if er, ok := err.(serverError); ok {
		code = httpErrors[er.code]
		action = er.action
		message = er.message
		errors = er.errors
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
func New(code int, action int, message string, errors ...any) error {
	var errs any = nil
	if len(errors) != 0 {
		errs = errors[0]
	}
	return serverError{
		code:    code,
		action:  action,
		message: message,
		errors:  errs,
	}
}
