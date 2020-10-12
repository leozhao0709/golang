package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/leozhao0709/golang/bookstore_user_api/app/users"
)

// RegisterRoute register route
func RegisterRoute(e *echo.Echo) {
	users.RegisterRoute(e)
}

// GenerateRouteJSON generate a json file containing all routes
func GenerateRouteJSON(e *echo.Echo) {
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("src/routes/routes.json", data, 0644)
}
