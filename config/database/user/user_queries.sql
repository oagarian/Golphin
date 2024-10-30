-- name: InsertUser :exec
INSERT INTO "users"(id, username, email, password)
VALUES ($1, $2, $3, $4);

-- name: FindUser :one
SELECT * FROM "users"
WHERE (username = $1 OR email = $2) AND password = $3;
