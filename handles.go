package main

import "net/http"

func testHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}