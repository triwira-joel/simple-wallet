package user

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *userRepository) CreateUser(c echo.Context, u model.CreateUserParams) error {
	_, err := r.DB.CreateUser(c.Request().Context(), u)
	if err != nil {
		log.Println("failed to create User", u, err.Error())
		return err
	}
	return nil
}
