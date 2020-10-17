package response

import "net/http"

// IErrorResponse common error response
type IErrorResponse interface {
	StatusCode() int
	Error() string
}

type errorResponse struct {
	ErrStatusCode int    `json:"-"`     // http statusCode
	ErrError      string `json:"error"` // error message
}

func (res errorResponse) StatusCode() int {
	return res.ErrStatusCode
}

func (res errorResponse) Error() string {
	return res.ErrError
}

// NewErrorResponse create a fully customize error response
func NewError(statusCode int, err error) IErrorResponse {
	return errorResponse{
		ErrStatusCode: statusCode,
		ErrError:      err.Error(),
	}
}

// BadRequest create a fully customize error response
func BadRequest(err error) IErrorResponse {
	return errorResponse{
		ErrStatusCode: http.StatusBadRequest,
		ErrError:      err.Error(),
	}
}

// InternalServerError create a fully customize error response
func InternalServerError(err error) IErrorResponse {
	return errorResponse{
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      err.Error(),
	}
}
