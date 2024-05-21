package user

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *userRepository) UpdateUser(c echo.Context, p model.UpdateUserParams) error {
	err := r.DB.UpdateUser(c.Request().Context(), p)
	if err != nil {
		log.Println("failed to update User", p, err.Error())
		return err
	}
	return nil
}
