package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/EmmanuelAllanMJ/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	feed_follow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error in following feed: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feed_follow))
}

func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting feed follows: %v", err))
		return
	}

	var feed_followsResponse []FeedFollow
	for _, feed_follow := range feed_follows {
		feed_followsResponse = append(feed_followsResponse, databaseFeedFollowToFeedFollow(feed_follow))
	}

	respondWithJson(w, 200, feed_followsResponse)
}

func (apiConfig *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID, err := uuid.Parse(chi.URLParam(r, "feedFollowID"))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid feedFollowID: %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedFollowID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}
