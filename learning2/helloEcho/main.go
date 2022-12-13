package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }), func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		fmt.Println("before call 1")
	// 		<-time.After(time.Second * 3)
	// 		next(c)
	// 		fmt.Println("after call 1")
	// 		return nil
	// 	}
	// }, func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		fmt.Println("before call 2")
	// 		<-time.After(time.Second * 3)
	// 		next(c)
	// 		fmt.Println("after call 2")
	// 		return nil
	// 	}
	// })

	e.GET("/", func(c echo.Context) error {
		fmt.Println("----start----")
		<-time.After(time.Second * 3)
		return c.JSON(http.StatusOK, &struct {
			Name string `json:"name"`
			Age  int8   `json:"age"`
		}{Name: "Lei", Age: 31})
	})

	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
