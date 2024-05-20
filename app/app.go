package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	userHandler "github.com/triwira-joel/simple-wallet/private/adapter/api/user"
	userRepo "github.com/triwira-joel/simple-wallet/private/adapter/repository/user"
	userService "github.com/triwira-joel/simple-wallet/private/service/user"

	_ "github.com/go-sql-driver/mysql"
	sqldb "github.com/triwira-joel/simple-wallet/sqldb"
)

// to init database for now, will remove later
func initDb() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3333)/simple_wallet?parseTime=true")
	if err != nil {
		return err
	}

	queries := sqldb.New(db)

	// list all authors
	authors, err := queries.GetUsers(ctx)
	if err != nil {
		return err
	}
	log.Println("Users:", authors)

	return nil
}
func main() {
	userRepo := userRepo.NewRepository()
	userService := userService.New(userRepo)
	userHandler := userHandler.NewHttpHandler(userService)

	if err := initDb(); err != nil {
		log.Fatal(err)
	}

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
