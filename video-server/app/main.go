package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/video-server/app/routes"
)

func main() {
	e := echo.New()

	e.Debug = true
	// log
	e.Debug = true
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("[${time_rfc3339}] ${level}")
	}

	// middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${method} ${uri} ${status}\n",
	}))

	// register route
	routes.RegisterRoute(e)

	// start server
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
