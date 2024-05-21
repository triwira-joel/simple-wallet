package account

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *accountRepository) UpdateAccount(c echo.Context, p model.UpdateAccountByUserIdParams) error {
	err := r.DB.UpdateAccountByUserId(c.Request().Context(), p)
	if err != nil {
		log.Println("failed to update Account", p, err.Error())
		return err
	}
	return nil
}
