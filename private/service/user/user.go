package user

import (
	"github.com/labstack/echo/v4"
	repo "github.com/triwira-joel/simple-wallet/private/adapter/repository/user"
	domain "github.com/triwira-joel/simple-wallet/private/domain/user"
)

type userService struct {
	repo repo.UserRepository
}

func NewUserService(
	repo repo.UserRepository,
) *userService {
	return &userService{
		repo,
	}
}

type UserService interface {
	GetUsers(c echo.Context) ([]domain.User, error)
}
