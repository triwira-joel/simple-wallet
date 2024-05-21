package transactionhistory

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *transactionHistoryRepository) UpdateTransactionHistoryByUserId(c echo.Context, p model.UpdateTransactionHistoryByUserIdParams) error {
	err := r.DB.UpdateTransactionHistoryByUserId(c.Request().Context(), p)
	if err != nil {
		log.Println("error update transaction history by user id", p, err.Error())
		return err
	}
	return nil
}
