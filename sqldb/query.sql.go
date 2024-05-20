// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :execresult
INSERT INTO accounts (
    user_id, balance
) VALUES (
    ?,?
)
`

type CreateAccountParams struct {
	UserID  sql.NullInt32
	Balance sql.NullInt64
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAccount, arg.UserID, arg.Balance)
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
    name, email, password, status, created_at, updated_at
) VALUES (
    ?,?,?,?,?,?
)
`

type CreateUserParams struct {
	Name      sql.NullString
	Email     sql.NullString
	Password  sql.NullString
	Status    sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const getAccount = `-- name: GetAccount :one
SELECT id, user_id, balance, created_at, updated_at FROM accounts
WHERE id = ? LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByUserId = `-- name: GetAccountByUserId :one
SELECT id, user_id, balance, created_at, updated_at FROM accounts
WHERE user_id = ? LIMIT 1
`

func (q *Queries) GetAccountByUserId(ctx context.Context, userID sql.NullInt32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUserId, userID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccounts = `-- name: GetAccounts :many
SELECT id, user_id, balance, created_at, updated_at FROM accounts
ORDER BY id
`

func (q *Queries) GetAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Balance,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, password, status, created_at, updated_at FROM users
WHERE id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, email, password, status, created_at, updated_at FROM users
ORDER BY id
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :exec
UPDATE accounts SET 
    balance=?,updated_at=?
WHERE id=? AND user_id=?
`

type UpdateAccountParams struct {
	Balance   sql.NullInt64
	UpdatedAt sql.NullTime
	ID        int32
	UserID    sql.NullInt32
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	_, err := q.db.ExecContext(ctx, updateAccount,
		arg.Balance,
		arg.UpdatedAt,
		arg.ID,
		arg.UserID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET 
    name=?,email=?,password=?,status=?,created_at=?,updated_at=?
WHERE id=?
`

type UpdateUserParams struct {
	Name      sql.NullString
	Email     sql.NullString
	Password  sql.NullString
	Status    sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	ID        int32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
