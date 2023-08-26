package main

import (
	"net/http"
)

func handleRediness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"status": "ok"})
}
