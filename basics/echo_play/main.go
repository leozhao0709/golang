package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/jellydator/validation"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Name  *string `json:"name" `
	Email *string `json:"email" `
	Age   *int    `json:"age,omitempty" `
}

type UserDto struct {
	Name  string
	Email string
	Age   int
}

type Validator interface {
	Validate() error
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Age, validation.NotNil, validation.Max(10)),
	)
}

func ValidateRequestBody[T Validator]() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			t := new(T)
			if err := c.Bind(t); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			if err := (*t).Validate(); err != nil {
				logger.Error("Validate request body failed: ", "err", err)
				return echo.NewHTTPError(http.StatusBadRequest, err)
			}

			c.Set("request_body", t)
			return next(c)
		}
	}
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	AddSource: true,
}))

func main() {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRequestID: true,
		LogRemoteIP:  true,
		LogHost:      true,
		LogMethod:    true,
		LogURI:       true,
		LogUserAgent: true,
		LogStatus:    true,
		LogLatency:   true,
		HandleError:  true, // forwards error to the global error handler, so it can decide appropriate status code
		BeforeNextFunc: func(c echo.Context) {
			// logger.Info("Request Start", "uri", c.Request().URL.Path)
			logger = logger.With("uri", c.Request().URL.Path)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
			logger.Info("Request end",
				slog.String("request_id", v.RequestID),
				slog.String("remote_ip", v.RemoteIP),
				slog.String("host", v.Host),
				slog.String("method", v.Method),
				slog.String("uri", v.URI),
				slog.String("user_agent", v.UserAgent),
				slog.Int("status", v.Status),
				slog.String("latency", v.Latency.String()),
			)
			return nil
		},
	}))

	e.POST("/users", func(c echo.Context) (err error) {
		u := c.Get("request_body").(*User)

		u_dto := &UserDto{}
		copier.Copy(u_dto, u)
		slog.Info("User: ", "u_dto", u_dto)

		return c.JSON(http.StatusOK, u)
	}, ValidateRequestBody[User]())

	e.Logger.Fatal(e.Start(":1323"))
}
