package account

import (
	"github.com/labstack/echo/v4"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

type accountRepository struct {
	DB *model.Queries
}

func NewAccountRepository(
	DB *model.Queries,
) *accountRepository {
	return &accountRepository{DB}
}

type AccountRepository interface {
	GetAccounts(c echo.Context) ([]model.Account, error)
	GetAccount(c echo.Context, id int32) (model.Account, error)
	CreateAccount(c echo.Context, u model.CreateAccountParams) error
	UpdateAccount(c echo.Context, p model.UpdateAccountByUserIdParams) error
}
