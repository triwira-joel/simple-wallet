package user

import (
	repo "github.com/triwira-joel/simple-wallet/private/port/repository/user"
)

type service struct {
	repo repo.UserRepository
}

func New(
	repo repo.UserRepository,
) *service {
	return &service{
		repo,
	}
}
