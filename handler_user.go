package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/EmmanuelAllanMJ/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiConfig.DB.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting posts for user: %v", err))
		return
	}

	respondWithJson(w, 200, databasePostsToPosts(posts))
}
