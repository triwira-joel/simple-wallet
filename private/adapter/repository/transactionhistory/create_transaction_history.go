package transactionhistory

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *transactionHistoryRepository) CreateTransactionHistory(c echo.Context, p model.CreateTransactionHistoryParams) error {
	_, err := r.DB.CreateTransactionHistory(c.Request().Context(), p)
	if err != nil {
		log.Println("error create transaction history", p, err.Error())
		return err
	}
	return nil
}
