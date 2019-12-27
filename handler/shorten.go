package handler

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v7"
)

//ShortenURL creates short version of given LongURL
func ShortenURL(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := client.Set("anish", "bhusal", 0).Err()
		if err != nil {
			panic(err)
		}
		val, err := client.Get("anish").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("value for key anish is:", val)

	}
}
