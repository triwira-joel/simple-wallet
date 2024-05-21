package transactionhistory

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

func (r *transactionHistoryRepository) GetTransactionHistoriesByAccountId(c echo.Context, id int32) ([]model.TransactionHistory, error) {
	tx_h, err := r.DB.GetTransactionHistoriesByAccountId(c.Request().Context(), sql.NullInt32{Int32: id})
	if err != nil {
		log.Println("error get transaction histories", err.Error())
		return tx_h, err
	}
	return tx_h, nil
}
