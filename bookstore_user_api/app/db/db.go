package db

import (
	"sync"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leozhao0709/golang/bookstore_user_api/ent"

	entsql "github.com/facebook/ent/dialect/sql"
)

var once sync.Once

// please add project ent dependency before
var client *ent.Client
var connectError error

// GetEntClient new ent client
func GetEntClient() (*ent.Client, error) {

	once.Do(func() {
		db, err := sql.Open("mysql", "[yourDbConnection]?parseTime=true")
		if err != nil {
			client = nil
			connectError = err
		}

		// Get the underlying sql.DB object of the driver.
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Hour)

		// Create an ent.Driver from `db`.
		drv := entsql.OpenDB("mysql", db)
		client = ent.NewClient(ent.Driver(drv)).Debug()
	})

	return client, connectError
}
