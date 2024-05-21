package user

import (
	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

type UserRepository interface {
	GetUsers(c echo.Context) ([]model.User, error)
}
