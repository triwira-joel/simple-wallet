package infra

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/triwira-joel/simple-wallet/constants"
)

type DBHandler struct {
	DB *sql.DB
}

func (d *DBHandler) ConnectDB() error {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3333)/simple_wallet?parseTime=true")
	if err != nil {
		log.Println(constants.ConnectDBFail + " | " + err.Error())
		return err
	}

	d.DB = db
	err = d.DB.Ping()
	if err != nil {
		log.Println(constants.ConnectDBFail, err.Error())
		return err
	}

	return nil
}

func (d *DBHandler) Close() {
	if err := d.DB.Close(); err != nil {
		log.Println(constants.ClosingDBReadFailed + " | " + err.Error())
	} else {
		log.Println(constants.ClosingDBReadSuccess)
	}
}
