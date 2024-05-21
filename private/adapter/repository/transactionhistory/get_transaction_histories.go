package transactionhistory

import (
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *transactionHistoryRepository) GetTransactionHistories(c echo.Context) ([]model.TransactionHistory, error) {
	tx_h, err := r.DB.GetTransactionHistories(c.Request().Context())
	if err != nil {
		log.Println("error get transaction histories", err.Error())
		return tx_h, err
	}
	return tx_h, nil
}
