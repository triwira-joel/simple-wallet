package user

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *userRepository) GetUsers(c echo.Context) ([]model.User, error) {
	users, err := r.DB.GetUsers(c.Request().Context())
	if err != nil {
		log.Println("error get users", err.Error())
		return users, err
	}
	return users, nil
}
