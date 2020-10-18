package response

import (
	"net/http"

	"github.com/leozhao0709/golang/bookstore_user_api/app/errcode"
)

// IErrorResponse common error response
type IErrorResponse interface {
	StatusCode() int
	ErrCode() int
	Error() string
}

type errorResponse struct {
	ErrStatusCode int    `json:"-"`          // http statusCode
	ErrErrorCode  int    `json:"error_code"` // business err code
	ErrError      string `json:"error"`      // error message
}

func (res *errorResponse) StatusCode() int {
	return res.ErrStatusCode
}

func (res *errorResponse) Error() string {
	return res.ErrError
}

func (res *errorResponse) ErrCode() int {
	return res.ErrErrorCode
}

// NewError create a fully customize error response
func NewError(statusCode int, errorCode int, err error) IErrorResponse {
	return &errorResponse{
		ErrStatusCode: statusCode,
		ErrError:      err.Error(),
		ErrErrorCode:  errorCode,
	}
}

// BadRequest create a fully customize error response
func BadRequest(errorCode int, err error) IErrorResponse {
	return &errorResponse{
		ErrStatusCode: http.StatusBadRequest,
		ErrErrorCode:  errorCode,
		ErrError:      err.Error(),
	}
}

// InternalServerError create a fully customize error response
func InternalServerError(err error) IErrorResponse {
	return &errorResponse{
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      err.Error(),
		ErrErrorCode:  errcode.InternalServerErr,
	}
}
