package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	mux.HandleFunc("/", testHandle)

	fmt.Println("Listening on port "+server.Addr)
	err := server.ListenAndServe()
	log.Fatal(err)
}