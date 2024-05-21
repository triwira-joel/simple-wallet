package user

import (
	repo "github.com/triwira-joel/simple-wallet/private/port/repository/user"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(
	repo repo.UserRepository,
) *UserService {
	return &UserService{
		repo,
	}
}
