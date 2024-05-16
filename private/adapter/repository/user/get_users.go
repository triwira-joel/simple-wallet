package user

import (
	"github.com/labstack/echo/v4"
	userModel "github.com/triwira-joel/simple-wallet/private/adapter/repository/db/user"
)

func (repo *repository) GetUsers(c echo.Context) ([]userModel.UserDB, error) {
	return users, nil
}
