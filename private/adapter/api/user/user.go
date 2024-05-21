package user

import userService "github.com/triwira-joel/simple-wallet/private/port/service/user"

type UserHTTPHandler struct {
	userService userService.UserService
}

func NewUserHTTPHandler(
	userService userService.UserService,
) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}
