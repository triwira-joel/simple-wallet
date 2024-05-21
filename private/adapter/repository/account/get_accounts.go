package account

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *accountRepository) GetAccounts(c echo.Context) ([]model.Account, error) {
	accounts, err := r.DB.GetAccounts(c.Request().Context())
	if err != nil {
		log.Println("failed to get Accounts", err.Error())
		return accounts, err
	}
	return accounts, nil
}
