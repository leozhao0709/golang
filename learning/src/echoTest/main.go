package main

import (
	"net/http"

	"github.com/gookit/color"
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
	log.SetHeader("${time_rfc3339} ${level}")

	e.GET("/ping", func(c echo.Context) error {
		// log.Info(color.Cyan.Renderln("this is a info logger"))
		// log.Info(color.BgCyan.Render("this is a info logger"))
		log.Info(color.RGB(30, 144, 255).Sprint("this is a info logger"))
		// log.Info(color.New(color.FgLightBlue, color.BgLightYellow).Renderln("this is a info logger"))
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
