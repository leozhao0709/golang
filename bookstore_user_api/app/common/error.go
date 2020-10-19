package common

import (
	"net/http"
)

// Common Error Code
const (
	RequestValidationErrCode = 4000
	InternalServerErrCode    = 5000
)

// ICommonError common error
type ICommonError interface {
	StatusCode() int
	ErrCode() int
	Error() string
}

type commonError struct {
	ErrStatusCode int    `json:"-"`          // http statusCode
	ErrErrorCode  int    `json:"error_code"` // business err code
	ErrError      string `json:"error"`      // error message
}

func (err *commonError) StatusCode() int {
	return err.ErrStatusCode
}

func (err *commonError) Error() string {
	return err.ErrError
}

func (err *commonError) ErrCode() int {
	return err.ErrErrorCode
}

// NewError create a fully customize error errponse
func NewError(statusCode int, errorCode int, err error) ICommonError {
	return &commonError{
		ErrStatusCode: statusCode,
		ErrError:      err.Error(),
		ErrErrorCode:  errorCode,
	}
}

// BadRequestError create a fully customize error errponse
func BadRequestError(errorCode int, err error) ICommonError {
	return &commonError{
		ErrStatusCode: http.StatusBadRequest,
		ErrErrorCode:  errorCode,
		ErrError:      err.Error(),
	}
}

// InternalServerError create a fully customize error errponse
func InternalServerError(err error) ICommonError {
	return &commonError{
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      err.Error(),
		ErrErrorCode:  InternalServerErrCode,
	}
}
