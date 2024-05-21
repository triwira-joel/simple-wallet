package transactionhistory

import (
	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

type transactionHistoryRepository struct {
	DB *model.Queries
}

func NewTransactionHistoryRepository(
	DB *model.Queries,
) *transactionHistoryRepository {
	return &transactionHistoryRepository{DB}
}

type TransactionHistory interface {
	GetTransactionHistories(c echo.Context) ([]model.TransactionHistory, error)
	GetTransactionHistoriesByUserId(c echo.Context, id int32) ([]model.TransactionHistory, error)
	GetTransactionHistoriesByAccountId(c echo.Context, id int32) ([]model.TransactionHistory, error)
	// GetLatestTransactionHistoryByUserId(c echo.Context, id int32) (model.TransactionHistory, error)
	// GetLatestTransactionHistoryByAccountId(c echo.Context, id int32) (model.TransactionHistory, error)
	CreateTransactionHistory(c echo.Context, p model.CreateTransactionHistoryParams) error
	UpdateTransactionHistoryByUserId(c echo.Context, p model.UpdateTransactionHistoryByUserIdParams) error
	UpdateTransactionHistoryByAccountId(c echo.Context, p model.UpdateTransactionHistoryByAccountIdParams) error
}
