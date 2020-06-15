package message

const (
	// LoginMsgType Login message type
	LoginMsgType = "LoginMsgType"
	// LoginMsgResType Login Message Result type
	LoginMsgResType = "LoginMsgResType"
)

// Message message type
type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// LoginMsg Login message
type LoginMsg struct {
	UserID   *string `json:"userId"`
	Password *string `json:"password"`
	Username *string `json:"username"`
}

// LoginMsgRes Login Message Result
type LoginMsgRes struct {
	Code  *int   `json:"code"`
	Error *error `json:"error"`
}
