-- name: GetAccountSessionByRefreshToken :one
SELECT *
FROM accounts_sessions
WHERE refresh_token = $1
LIMIT 1;

-- name: CreateAccountSession :exec
INSERT INTO accounts_sessions (
        account_id,
        refresh_token,
        user_agent,
        ip_address,
        issued_at,
        expires_at
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: DeleteAccountSession :exec
DELETE FROM accounts_sessions
WHERE refresh_token = $1;
