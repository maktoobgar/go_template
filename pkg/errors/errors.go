package errors

import (
	"net/http"

	"github.com/maktoobgar/go_template/pkg/errors/messages"
)

type (
	serverError struct {
		code    int
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

// Returns httpErrorCode and message of it
func HttpError(e error) (string, int) {
	err, ok := e.(serverError)
	if !ok {
		return messages.ErrorGeneral, httpErrors[UnexpectedStatus]
	}

	httpCode, ok := httpErrors[err.code]
	if !ok {
		return messages.ErrorGeneralNotFound, httpErrors[InvalidStatus]
	}

	return err.message, httpCode
}

// Creates a new error
func New(code int, message string) error {
	return serverError{
		code:    code,
		message: message,
	}
}
