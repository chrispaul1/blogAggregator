-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMPTZ NOT NULL,
    feed_id UUID NOT NULL
);

