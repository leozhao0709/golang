package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func ping(c echo.Context) error {
	log.Debug("this is a info logger")
	result := map[string]string{
		"message": "pong",
	}
	return c.JSON(http.StatusOK, result)
}
