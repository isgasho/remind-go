package models

import (
	"database/sql"
	"fmt"
	"log"
	"remind-go/config"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	configure := config.LoadConfig()
	driver := configure.Db.Driver
	source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true",
		configure.Db.User, configure.Db.Password,
		configure.Db.Address, configure.Db.Database)
	Db, err = sql.Open(driver, source)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
