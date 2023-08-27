package main

import (
	"net/http"

	"github.com/EmmanuelAllanMJ/rssagg/internal/auth"
	"github.com/EmmanuelAllanMJ/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, "Authentication error")
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, "Error getting user")
			return
		}

		handler(w, r, user)
	}
}
