package handler

import (
	"encoding/json"
	"net/http"
)

//SendResponseToClient sends application/json response to client
func SendResponseToClient(w http.ResponseWriter, r *http.Request, response interface{}, statusCode int) {
	//send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
	return
}
