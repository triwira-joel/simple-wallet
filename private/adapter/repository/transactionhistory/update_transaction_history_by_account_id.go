package transactionhistory

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *transactionHistoryRepository) UpdateTransactionHistoryByAccountId(c echo.Context, p model.UpdateTransactionHistoryByAccountIdParams) error {
	err := r.DB.UpdateTransactionHistoryByAccountId(c.Request().Context(), p)
	if err != nil {
		log.Println("error update transaction history by user id", p, err.Error())
		return err
	}
	return nil
}
