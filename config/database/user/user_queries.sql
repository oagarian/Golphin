-- name: InsertUser :exec
INSERT INTO user (id, username, email, password)
VALUES ($1, $2, $3, $4);

-- name: FindUser :one
SELECT * FROM user
WHERE (username = $1 OR email = $2) AND password = $3;
