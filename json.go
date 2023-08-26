package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5xx error %d: %s", code, msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errorResponse{Error: msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json response : %v", err)
		w.WriteHeader(http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

// this is a helper function to respond with an error message in json format to the client
