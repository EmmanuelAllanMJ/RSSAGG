-- name: CreatePost :one
INSERT INTO posts (id, title, description, url, published_at, created_at, updated_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, title, description, url, published_at, created_at, updated_at, feed_id;

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    title TEXT,
    description TEXT,
    url TEXT NOT NULL ,
    published_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- name: GetPostForUser :many
SELECT posts.* FROM posts
JOIN feed_follow ON posts.feed_id = feed_follow.feed_id
WHERE feed_follow.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;