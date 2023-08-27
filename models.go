package main

import (
	"time"

	"github.com/EmmanuelAllanMJ/rssagg/internal/database"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID.String(),
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    string    `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID.String(),
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID.String(),
	}
}

type FeedFollow struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`
	FeedID    string    `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID.String(),
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID.String(),
		FeedID:    dbFeedFollow.FeedID.String(),
	}
}

type Post struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Url         string     `json:"url"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      string     `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {
	var publishedAt *time.Time
	if dbPost.PublishedAt.Valid {
		publishedAt = &dbPost.PublishedAt.Time
	}
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	var createdAt time.Time
	if dbPost.CreatedAt.Valid {
		createdAt = dbPost.CreatedAt.Time
	}
	var updatedAt time.Time
	if dbPost.UpdatedAt.Valid {
		updatedAt = dbPost.UpdatedAt.Time
	}

	return Post{
		ID:          dbPost.ID.String(),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Title:       dbPost.Title.String,
		Description: description,
		Url:         dbPost.Url,
		PublishedAt: publishedAt,
		FeedID:      dbPost.FeedID.String(),
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := make([]Post, len(dbPosts))
	for i, dbPost := range dbPosts {
		posts[i] = databasePostToPost(dbPost)
	}
	return posts
}
