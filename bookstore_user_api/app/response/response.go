package response

import "net/http"

// IResponse common response interface
type IResponse interface {
	StatusCode() int
	Data() interface{}
	Message() string
}

type response struct {
	ResStatusCode int         `json:"-"` // http statusCode
	ResData       interface{} `json:"data"`
	ResMessage    string      `json:"message"`
}

func (res response) StatusCode() int {
	return res.ResStatusCode
}
func (res response) Data() interface{} {
	return res.ResData
}
func (res response) Message() string {
	return res.ResMessage
}

// SuccessResponse create a simple success response
func SuccessResponse(data interface{}) IResponse {
	return response{
		ResStatusCode: http.StatusOK,
		ResMessage:    "success",
		ResData:       data,
	}
}

// NewResponse create a fully customize response
func NewResponse(statusCode int, message string, data interface{}) IResponse {
	return response{
		ResStatusCode: statusCode,
		ResMessage:    message,
		ResData:       data,
	}
}
