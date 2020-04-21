package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql" //init
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open(`mysql`, `goweb_dev:SDTa123vrakdf@tcp(115.236.176.108:10000)/goweb_dev?charset=utf8&parseTime=true`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	Db = db
}
