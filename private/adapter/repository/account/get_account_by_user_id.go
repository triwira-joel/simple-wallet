package account

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *accountRepository) CreateAccount(c echo.Context, p model.CreateAccountParams) error {
	_, err := r.DB.CreateAccount(c.Request().Context(), p)
	if err != nil {
		log.Println("failed to create Account", p, err.Error())
		return err
	}
	return nil
}
