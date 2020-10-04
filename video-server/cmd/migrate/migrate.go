package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/video-server/ent"
	"github.com/leozhao0709/golang/video-server/ent/migrate"
)

func main() {
	client, err := ent.Open("mysql", "lzhao:12345@tcp(localhost:3306)/video_server")
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
