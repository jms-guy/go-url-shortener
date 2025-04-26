package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jms-guy/go-url-shortener/shortener"
	"github.com/jms-guy/go-url-shortener/store"
)

type APIConfig struct {
	Db store.StorageService
}

type urlRequest struct {
	Url string `json:"url"`
}

func (a *APIConfig) redirectHandle(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("shortUrl")

	if shortUrl == "" {
		log.Println("Missing url data")
		respondWithError(w, 400, "Missing url data")
		return
	}

	redirect, err := store.GetInitialUrl(shortUrl)
	if err != nil {
		log.Printf("Error retrieving initial url from database: %s", err)
		respondWithError(w, 500, "Error getting url from database")
		return
	}

	http.Redirect(w, r, redirect, http.StatusFound)
}

func (a *APIConfig) shortenHandle(w http.ResponseWriter, r *http.Request) {
	request := urlRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		log.Printf("Error decoding input string: %s", err)
		respondWithError(w, 400, "Invalid JSON payload")
		return
	}

	intialUrl := request.Url
	shortVersion := shortener.GenerateShortLink(intialUrl)
	if shortVersion == "" {
		log.Printf("Error in url generation")
		respondWithError(w, 500, "Error in url generation")
		return
	}

	newUrl := shortener.SetNewUrl(shortVersion)
	err = store.SaveUrlMap(newUrl, intialUrl)
	if err != nil {
		log.Printf("Error saving data to redis database: %s", err)
		respondWithError(w, 500, "Error accessing server database")
		return
	}

	respondWithJSON(w, 200, urlRequest{Url: newUrl})
}