package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
)

// RegisterRoute register route
func RegisterRoute(e *echo.Echo) {

}

// GenerateRouteJSON generate a json file containing all routes
func GenerateRouteJSON(e *echo.Echo) {
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("src/routes/routes.json", data, 0644)
}
