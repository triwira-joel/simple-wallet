package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/triwira-joel/simple-wallet/constants"
	"github.com/triwira-joel/simple-wallet/private/config"
	sqldb "github.com/triwira-joel/simple-wallet/sqldb"
)

type DBHandler struct {
	DB      *sql.DB
	Queries *sqldb.Queries
}

func NewDBHandler(db *sql.DB) *DBHandler {
	return &DBHandler{
		DB:      db,
		Queries: sqldb.New(db),
	}
}

func (d *DBHandler) ConnectDB(cnf *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cnf.Database.User, cnf.Database.Password,
		cnf.Database.Host,
		cnf.Database.Port,
		cnf.Database.Name,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(constants.ConnectDBFail + " | " + err.Error())
	}

	d.DB = db
	err = d.DB.Ping()
	if err != nil {
		log.Fatal(constants.ConnectDBFail, err.Error())
	}

	return db
}

func (d *DBHandler) Close() {
	if err := d.DB.Close(); err != nil {
		log.Println(constants.ClosingDBReadFailed + " | " + err.Error())
	} else {
		log.Println(constants.ClosingDBReadSuccess)
	}
}
