package mydb

import (
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "lzhao:12345@tcp(127.0.0.1:3307)/fileserver?charset=utf8")
	log.Info("......db....", db)
	if err != nil {
		log.Panic("db connect err", err)
	}
	db.SetMaxOpenConns(1000)
}

// GetDB return the db connection
func GetDB() *sqlx.DB {
	log.Info("......db....", db)
	return db
}
