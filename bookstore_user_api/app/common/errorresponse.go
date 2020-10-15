package common

// IErrorResponse common error response
type IErrorResponse interface {
	StatusCode() int
	Code() int
	Error() string
}

type errorResponse struct {
	ErrStatusCode int    `json:"-"`     // http statusCode
	ErrError      string `json:"error"` // error message
	ErrCode       int    `json:"code"`  // your business error code
}

func (res errorResponse) StatusCode() int {
	return res.ErrStatusCode
}
func (res errorResponse) Code() int {
	return res.ErrCode
}
func (res errorResponse) Error() string {
	return res.ErrError
}

// NewErrorResponse create a fully customize error response
func NewErrorResponse(statusCode int, err error, code int) IErrorResponse {
	return errorResponse{
		ErrStatusCode: statusCode,
		ErrError:      err.Error(),
		ErrCode:       code,
	}
}
