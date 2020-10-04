package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/leozhao0709/golang/video-server/app/user"
)

// RegisterRoute register route
func RegisterRoute(e *echo.Echo) {
	user.RegisterRoute(e)
}

// GenerateRouteJSON generate a json file containing all routes
func GenerateRouteJSON(e *echo.Echo) {
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("src/routes/routes.json", data, 0644)
}
