package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (hdl *UserHTTPHandler) GetUsers(c echo.Context) error {
	users, err := hdl.userService.GetUsers(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, users)
}
