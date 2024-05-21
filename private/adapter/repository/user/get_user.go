package user

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *userRepository) GetUser(c echo.Context, id int32) (model.User, error) {
	user, err := r.DB.GetUser(c.Request().Context(), id)
	if err != nil {
		log.Println("failed to get User by id", id, err.Error())
		return user, err
	}
	return user, nil
}
