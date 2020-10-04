package common

// ErrorResponse error response
type ErrorResponse struct {
	ErrorMsg  string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	// BadRequest bad request
	BadRequest = ErrorResponse{ErrorMsg: "Bad Request", ErrorCode: "001"}
	// AuthError authentication error
	AuthError = ErrorResponse{ErrorMsg: "User authentication failed", ErrorCode: "002"}
)
