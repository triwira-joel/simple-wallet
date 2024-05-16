package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	userHandler "github.com/triwira-joel/simple-wallet/private/adapter/api/user"
	userRepo "github.com/triwira-joel/simple-wallet/private/adapter/repository/user"
	userService "github.com/triwira-joel/simple-wallet/private/service/user"
)

func main() {
	userRepo := userRepo.NewRepository()
	userService := userService.New(userRepo)
	userHandler := userHandler.NewHttpHandler(userService)
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", userHandler.GetUsers)
	e.Logger.Fatal(e.Start(":1323"))
}
