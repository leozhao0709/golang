package db

import (
	"sync"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/leozhao0709/golang/video-server/ent"
)

var once sync.Once

var client *ent.Client
var connectError error

// GetEntClient new ent client
func GetEntClient() (*ent.Client, error) {

	once.Do(func() {
		db, err := sql.Open("mysql", "lzhao:12345@tcp(localhost:3306)/video_server?parseTime=true")
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
