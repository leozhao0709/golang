package handlererror

import "net/http"

// HandleError http handler error
type HandleError struct {
	StatusCode int
	Err        error
}

// CreateError create a customer error with code
func CreateError(code int, err error) *HandleError {
	return &HandleError{StatusCode: code, Err: err}
}

// InternalServerError Internal Server Error
func InternalServerError(err error) *HandleError {
	return CreateError(http.StatusInternalServerError, err)
}

// ForbiddenError Forbidden Error
func ForbiddenError(err error) *HandleError {
	return CreateError(http.StatusForbidden, err)
}

// MethodNotAllowedError Method Not Allowed Error
func MethodNotAllowedError(err error) *HandleError {
	return CreateError(http.StatusMethodNotAllowed, err)
}

// NotFoundError 404 not found error
func NotFoundError(err error) *HandleError {
	return CreateError(http.StatusNotFound, err)
}
