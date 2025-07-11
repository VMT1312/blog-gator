-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT not NULL,
    url TEXT UNIQUE not NULL,
    user_id UUID not NULL,
    foreign key (user_id)
    references users(id) on DELETE cascade
);

-- +goose Down
DROP TABLE feeds;