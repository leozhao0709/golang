package common

import "net/http"

// IResponse common response interface
type IResponse interface {
	StatusCode() int
	Code() int
	Data() interface{}
	Message() string
}

type response struct {
	ResStatusCode int         `json:"-"`    // http statusCode
	ResCode       int         `json:"code"` // your business response code
	ResData       interface{} `json:"data"`
	ResMessage    string      `json:"message"`
}

func (res response) StatusCode() int {
	return res.ResStatusCode
}
func (res response) Code() int {
	return res.ResCode
}
func (res response) Data() interface{} {
	return res.ResData
}
func (res response) Message() string {
	return res.ResMessage
}

// SuccessResponse create a simple success response
func SuccessResponse(code int, data interface{}) IResponse {
	return response{
		ResStatusCode: http.StatusOK,
		ResMessage:    "success",
		ResCode:       code,
		ResData:       data,
	}
}

// NewResponse create a fully customize response
func NewResponse(statusCode int, message string, code int, data interface{}) IResponse {
	return response{
		ResStatusCode: statusCode,
		ResMessage:    message,
		ResCode:       code,
		ResData:       data,
	}
}
