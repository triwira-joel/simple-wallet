package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/triwira-joel/simple-wallet/infra"
	userHandler "github.com/triwira-joel/simple-wallet/private/adapter/api/user"
	repo "github.com/triwira-joel/simple-wallet/private/adapter/repository/user"
	"github.com/triwira-joel/simple-wallet/private/config"
	userService "github.com/triwira-joel/simple-wallet/private/service/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cnf := config.Get()
	dbStruct := infra.DBHandler{}
	db := dbStruct.ConnectDB(cnf)

	dbHandler := infra.NewDBHandler(db)
	repo := repo.NewUserRepository(dbHandler.Queries)
	userService := userService.NewUserService(repo)
	userHandler := userHandler.NewUserHTTPHandler(userService)

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
