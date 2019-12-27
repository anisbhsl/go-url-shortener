package handler

import (
	"go-url-shortener/model"
	"log"
	"math/rand"
	"net/http"

	"github.com/go-redis/redis/v7"
)

//ShortenURL creates short version of given LongURL
func ShortenURL(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var url model.URL
		queryURL := r.FormValue("url")
		if queryURL == "" {
			var errorResponse model.ErrorResponse
			errorResponse.Message = "URL Param Empty"
			errorResponse.MessageDescription = "Long URL should be provided as query param to generate shortened URL"

			SendResponseToClient(w, r, errorResponse, 400)
			return
		}

		url.LongURL = queryURL

		//check if URL already exists in db
		val, err := client.Get(queryURL).Result()
		if err == redis.Nil {
			//key doesn't exist
			log.Println("[[handler/shorten.go]] ShortURL doesn't exist!")
			shortURL := generateHash(queryURL)
			url.ShortURL = shortURL

			insertionErr := client.Set(url.LongURL, url.ShortURL, 0).Err()
			if insertionErr != nil {
				var errorResponse model.ErrorResponse
				errorResponse.Message = "Oops Error Occured! Please try again!"
				//send response
				SendResponseToClient(w, r, errorResponse, 400)
				return
			}

			reverseInsertionErr := client.Set(url.ShortURL, url.LongURL, 0).Err()
			if reverseInsertionErr != nil {
				var errorResponse model.ErrorResponse
				errorResponse.Message = "Oops Error Occured! Please try again!"
				//send response
				SendResponseToClient(w, r, errorResponse, 400)
				return
			}

			//send response
			SendResponseToClient(w, r, url, 200)
			return
		} else if err != nil {
			log.Println("[[handler/shorten.go]] Error: ", err)
			var errorResponse model.ErrorResponse
			errorResponse.Message = "Oops Error Occured! Please try again!"
			//send response
			SendResponseToClient(w, r, errorResponse, 400)
			return

		} else {
			log.Println("[[handler/shorten.go]] ShortURL already exists: ", val)

			//if it already exists
			var newurl model.URL
			newurl.LongURL = queryURL
			newurl.ShortURL = val
			//send response
			SendResponseToClient(w, r, newurl, 200)
			return
		}
	}
}

//generateHash generates a random 6 digit hash based on cropus
func generateHash(url string) string {
	var corpus = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	hash := make([]byte, 6)
	for i := range hash {
		hash[i] = corpus[rand.Intn(len(corpus))]
	}
	return string(hash)
}
