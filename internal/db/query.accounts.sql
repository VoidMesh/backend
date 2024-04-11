-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1;

-- name: CheckAccountExistsByEmail :one
SELECT EXISTS (
        SELECT 1
        FROM accounts
        WHERE email = $1
    );

-- name: GetAccountByEmail :one
SELECT *
FROM accounts
WHERE email = $1
LIMIT 1;

-- name: ListAccounts :many
SELECT *
FROM accounts
ORDER BY created_at DESC;

-- name: CreateAccount :one
INSERT INTO accounts (email, password_hash)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateAccount :exec
UPDATE accounts
SET email = $2,
    password_hash = $3
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
