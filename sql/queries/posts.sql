-- name: CreatePost :exec
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1, 
    $2, 
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);

-- name: GetPost :many
SELECT * FROM posts
ORDER BY updated_at DESC
LIMIT $1;