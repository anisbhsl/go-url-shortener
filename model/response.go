package model

//Response is encoded as json to send application/json message to client
type Response struct {
	Message     string `json:"msg"`
	Description string `json:"desc,omitempty"`
}

//ErrorResponse is for custom error response to client
type ErrorResponse struct {
	Message            string `json:"err"`
	MessageDescription string `json:"err_desc,omitempty"`
}
