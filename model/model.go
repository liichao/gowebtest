package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql" //init
	"github.com/jmoiron/sqlx"
)

// Db 定义一个Db操作
var Db *sqlx.DB

func init() {
	db, err := sqlx.Open(`mysql`, `goweb_dev:SDTa123vrakdf@tcp(115.236.176.108:10000)/goweb_dev?charset=utf8&parseTime=true`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// 检查数据库连接是否正常
	if err = db.Ping(); err != nil {
		log.Fatalln(err.Error())
	}
	Db = db
}
