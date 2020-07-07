package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	// log
	e.Debug = true
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("[${time_rfc3339}] ${level}")
	}

	// middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${method} ${uri} ${status}\n",
	}))

	// generate routes json file
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("src/routes/routes.json", data, 0644)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
