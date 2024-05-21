package account

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *accountRepository) GetAccount(c echo.Context, userId int32) (model.Account, error) {
	account, err := r.DB.GetAccount(c.Request().Context(), userId)
	if err != nil {
		log.Println("failed to get Account by user_id", userId, err.Error())
		return account, err
	}
	return account, nil
}
