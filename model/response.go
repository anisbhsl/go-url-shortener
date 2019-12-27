package model

//Response is encoded as json to send application/json message to client
type Response struct {
	Message string `json:"msg"`
}
