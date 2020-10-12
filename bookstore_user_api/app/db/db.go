package db

import (
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/bookstore_user_api/config/dbconfig"
	"github.com/leozhao0709/golang/bookstore_user_api/ent"

	"github.com/facebook/ent/dialect/sql"
)

var once sync.Once

// please add project ent dependency before
var client *ent.Client

// GetEntClient new ent client
func GetEntClient() *ent.Client {

	once.Do(func() {
		config := dbconfig.GetConfig()
		datasource := config.GetDataSource()
		dbDriver := config.GetDriver()
		maxIdleConns := config.GetMaxIdleConns()
		maxOpenConns := config.GetMaxOpenConns()
		connMaxLifetime := config.GetConnMaxLifetime()

		drv, err := sql.Open(dbDriver, datasource)
		if err != nil {
			log.Fatalf("Get ent client error: %v", err)
		}

		// Get the underlying sql.DB object of the driver.
		db := drv.DB()
		db.SetMaxIdleConns(maxIdleConns)
		db.SetMaxOpenConns(maxOpenConns)
		db.SetConnMaxLifetime(time.Second * time.Duration(connMaxLifetime))

		client = ent.NewClient(ent.Driver(drv)).Debug()
	})

	return client
}
