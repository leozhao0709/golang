package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/bookstore_user_api/app/routes"
	"github.com/leozhao0709/golang/bookstore_user_api/env"
)

func main() {

	var environment = env.GetCurrentEnv()

	e := echo.New()

	if environment != "prod" {
		e.Debug = true
		log.SetLevel(log.DEBUG)
	}

	log.SetHeader("[${time_rfc3339}] ${level} ${long_file}:${line}")

	// middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${method} ${uri} ${status}\n",
	}))

	// register route (import your routes)
	routes.RegisterRoute(e)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// start server
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
