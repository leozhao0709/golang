package main

import (
	"net/http"

	// "github.com/gookit/color"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${method} ${status} uri=${uri} ${latency_human}\n",
	}))

	log.SetLevel(log.DEBUG)
	log.SetHeader("[${time_rfc3339}] ${level}")

	e.GET("/ping", ping)

	e.GET("/", func(c echo.Context) error {
		// <-time.After(time.Second * 5)
		return c.String(http.StatusOK, "Hello Golang")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
