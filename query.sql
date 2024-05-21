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

-- name: UpdateAccountByUserId :exec
UPDATE accounts SET 
    balance=?,updated_at=?
WHERE user_id=?;


-- name: GetTransactionHistory :one
SELECT * FROM transaction_history
WHERE id = ? LIMIT 1;

-- name: GetLatestTransactionHistoryByUserId :one
SELECT * FROM transaction_history
WHERE user_id = ? ORDER BY timestamp DESC LIMIT 1;

-- name: GetLatestTransactionHistoryByAccountId :one
SELECT * FROM transaction_history
WHERE account_id = ? ORDER BY timestamp DESC LIMIT 1;

-- name: GetTransactionHistories :many
SELECT * FROM transaction_history
ORDER BY id;

-- name: GetTransactionHistoriesByUserId :many
SELECT * FROM transaction_history
WHERE user_id = ?;

-- name: GetTransactionHistoriesByAccountId :many
SELECT * FROM transaction_history
WHERE account_id = ?;

-- name: CreateTransactionHistory :execresult
INSERT INTO transaction_history (
    transaction_type, amount, user_id, account_id, timestamp, created_at, updated_at
) VALUES (
    ?,?,?,?,?,?,?
);

-- name: UpdateTransactionHistoryByUserId :exec
UPDATE transaction_history SET 
    transaction_type=?,amount=?,user_id=?,account_id=?,timestamp=?,updated_at=?
WHERE user_id=?;

-- name: UpdateTransactionHistoryByAccountId :exec
UPDATE transaction_history SET 
    transaction_type=?,amount=?,user_id=?,account_id=?,timestamp=?,updated_at=?
WHERE account_id=?;