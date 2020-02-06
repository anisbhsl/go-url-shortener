package handler

import (
	"go-url-shortener/model"
	"net/http"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
)

//Redirect receives id as queryparam and redirect to original URL
func Redirect(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		if id == "" {
			var errorResponse model.ErrorResponse
			errorResponse.Message = "Bad Request: Empty short URL"
			errorResponse.MessageDescription = "use /shorten?url=<yourURL> to generate new short URL"
			//send response
			SendResponseToClient(w, r, errorResponse, 400)
			return
		}
		longURL, err := client.Get(id).Result()
		if err != nil {
			var errorResponse model.ErrorResponse
			errorResponse.Message = "URL has been removed from service!"
			errorResponse.MessageDescription = "use /shorten?url=<yourURL> to generate new short URL"
			//send response
			SendResponseToClient(w, r, errorResponse, 404)
			return
		}

		var response model.URL
		response.ShortURL = id
		response.LongURL = longURL
		//send response
		// SendResponseToClient(w, r, response, 200)
		http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
		return

	}
}
