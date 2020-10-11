package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/bookstore_user_api/ent"
	"github.com/leozhao0709/golang/bookstore_user_api/ent/migrate"
)

// add project ent/migrate dependency before
func main() {
	client, err := ent.Open("mysql", "lzhao:12345@tcp(127.0.0.1:3306)/bookstore_user?parseTime=true")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithFixture(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
