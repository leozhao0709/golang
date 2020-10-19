package common

import "net/http"

// IResponse common response interface
type IResponse interface {
	StatusCode() int
	Data() interface{}
}

type response struct {
	ResStatusCode int
	ResData       interface{}
}

func (res *response) StatusCode() int {
	return res.ResStatusCode
}
func (res *response) Data() interface{} {
	return res.ResData
}

// SuccessResponseWithCode create a simple success response with code
func SuccessResponseWithCode(code int, data interface{}) IResponse {
	return &response{
		ResStatusCode: http.StatusOK,
		ResData:       data,
	}
}

// SuccessResponse create a simple success response
func SuccessResponse(data interface{}) IResponse {
	return &response{
		ResStatusCode: http.StatusOK,
		ResData:       data,
	}
}

// NewResponse create a fully customize response
func NewResponse(statusCode int, data interface{}) IResponse {
	return &response{
		ResStatusCode: statusCode,
		ResData:       data,
	}
}
