package user

import (
	"github.com/labstack/echo/v4"
	db "github.com/triwira-joel/simple-wallet/sqldb"
	model "github.com/triwira-joel/simple-wallet/sqldb"
)

type userRepository struct {
	DB *db.Queries
}

func NewUserRepository(
	DB *db.Queries,
) *userRepository {
	return &userRepository{DB}
}

type UserRepository interface {
	GetUsers(c echo.Context) ([]model.User, error)
	GetUser(c echo.Context, id int32) (model.User, error)
	CreateUser(c echo.Context, u model.CreateUserParams) error
	UpdateUser(c echo.Context, p model.UpdateUserParams) error
}
