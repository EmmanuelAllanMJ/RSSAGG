// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: posts.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, title, description, url, published_at, created_at, updated_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, title, description, url, published_at, created_at, updated_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	Title       sql.NullString
	Description sql.NullString
	Url         string
	PublishedAt sql.NullTime
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	FeedID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Url,
		arg.PublishedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Url,
		&i.PublishedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedID,
	)
	return i, err
}

const getPostForUser = `-- name: GetPostForUser :many
SELECT posts.id, posts.title, posts.description, posts.url, posts.published_at, posts.created_at, posts.updated_at, posts.feed_id FROM posts
JOIN feed_follow ON posts.feed_id = feed_follow.feed_id
WHERE feed_follow.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2
`

type GetPostForUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

func (q *Queries) GetPostForUser(ctx context.Context, arg GetPostForUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Url,
			&i.PublishedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
