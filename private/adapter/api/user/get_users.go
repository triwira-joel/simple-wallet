package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	userService "github.com/triwira-joel/simple-wallet/private/port/service/user"
)

type HTTPHandler struct {
	userService userService.UserService
}

func NewHttpHandler(
	userService userService.UserService,
) *HTTPHandler {
	return &HTTPHandler{
		userService: userService,
	}
}

func (hdl *HTTPHandler) GetUsers(c echo.Context) error {
	users, err := hdl.userService.GetUsers(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, users)
}
