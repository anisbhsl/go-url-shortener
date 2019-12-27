package handler

import (
	"encoding/json"
	"go-url-shortener/model"
	"log"
	"net/http"
)

//IndexPage is URL Shortner API homepage handler
func IndexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[[handler/index.go]] Incoming Request")
		var response model.Response
		response.Message = "URL Shortener API"

		//send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
