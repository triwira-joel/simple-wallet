package user

import (
	db "github.com/triwira-joel/simple-wallet/sqldb"
)

type UserRepository struct {
	DB *db.Queries
}

func NewUserRepository(
	DB *db.Queries,
) *UserRepository {
	return &UserRepository{DB}
}
