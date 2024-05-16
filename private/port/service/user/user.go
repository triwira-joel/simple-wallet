package user

import (
	"github.com/labstack/echo/v4"
	domain "github.com/triwira-joel/simple-wallet/private/domain/user"
)

type UserService interface {
	GetUsers(c echo.Context) ([]domain.User, error)
}
