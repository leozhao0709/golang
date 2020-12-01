package app

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/bookstore_user_api/app/errorhandler"
	"github.com/leozhao0709/golang/bookstore_user_api/app/routes"
	"github.com/leozhao0709/golang/bookstore_user_api/env"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// StartApplication app entry point
func StartApplication() {
	var environment = env.GetCurrentEnv()

	t := &Template{
		templates: template.Must(template.ParseGlob("app/public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	if environment != "prod" {
		e.Debug = true
		log.SetLevel(log.DEBUG)
	}

	log.SetHeader("[${time_rfc3339}] ${level} ${long_file}:${line}")

	// middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${method} ${uri} ${status}\n",
	}))

	e.HTTPErrorHandler = errorhandler.RestfulHandler

	// register route (import your routes)
	routes.RegisterRoute(e)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// start server
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
