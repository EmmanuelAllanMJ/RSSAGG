-- +goose Up
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

-- +goose Down
DROP TABLE posts;