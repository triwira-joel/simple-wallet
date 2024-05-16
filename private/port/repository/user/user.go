package user

import (
	"github.com/labstack/echo/v4"
	userModel "github.com/triwira-joel/simple-wallet/private/adapter/repository/db/user"
)

type UserRepository interface {
	GetUsers(c echo.Context) ([]userModel.UserDB, error)
}
