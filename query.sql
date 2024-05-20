-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (
    name, email, password, status, created_at, updated_at
) VALUES (
    ?,?,?,?,?,?
);

-- name: UpdateUser :exec
UPDATE users SET 
    name=?,email=?,password=?,status=?,created_at=?,updated_at=?
WHERE id=?;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1;

-- name: GetAccountByUserId :one
SELECT * FROM accounts
WHERE user_id = ? LIMIT 1;

-- name: GetAccounts :many
SELECT * FROM accounts
ORDER BY id;

-- name: CreateAccount :execresult
INSERT INTO accounts (
    user_id, balance
) VALUES (
    ?,?
);

-- name: UpdateAccount :exec
UPDATE accounts SET 
    balance=?,updated_at=?
WHERE id=? AND user_id=?;