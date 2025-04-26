package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jms-guy/go-url-shortener/store"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	addr := os.Getenv("PORT")

	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr: ":" + addr,
	}

	api := APIConfig{
		Db: *store.InitializeStore(),
	}

	mux.HandleFunc("/", testHandle)
	mux.HandleFunc("POST /api/urls", api.shortenHandle)
	mux.HandleFunc("GET /{shortUrl}", api.redirectHandle)

	fmt.Println("Listening on port ", server.Addr)
	err := server.ListenAndServe()
	log.Fatal(err)
}