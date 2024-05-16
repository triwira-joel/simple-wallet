package user

import (
	"github.com/labstack/echo/v4"
	domain "github.com/triwira-joel/simple-wallet/private/domain/user"
)

func (srv *service) GetUsers(c echo.Context) ([]domain.User, error) {
	res := []domain.User{}
	users, err := srv.repo.GetUsers(c)
	if err != nil {
		return []domain.User{}, err
	}
	for _, user := range users {
		res = append(res, domain.User{Id: user.Id, Name: user.Name})
	}
	return res, nil
}
