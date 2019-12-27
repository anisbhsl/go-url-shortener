package main

import (
	"go-url-shortener/handler"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
)

const (
	serverPort string = "8000"
	redisPort  string = "6379"
)

//init is invoked before main
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env found!")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = serverPort
	}

	//create a new Redis client
	client, err := newRedisClient()
	if err != nil {
		panic(err)
	}

	log.Println("[[server.go]]Starting Server")
	router := mux.NewRouter()
	router.HandleFunc("/", handler.IndexPage()).Methods("GET")
	router.HandleFunc("/shorten", handler.ShortenURL(client)).Methods("GET")
	// router.HandleFunc("/{id}", handler.RedirectHandler()).Methods("GET")

	log.Println("[[server.go]] Listening at PORT:", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, router))
}

//newRedisClient creates a new client and returns *redis.Client,error
func newRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + redisPort,
		Password: "helloworld",
		DB:       0, //use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, err
}
