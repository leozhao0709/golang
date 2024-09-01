package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/jellydator/validation"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JobReq struct {
	Title *string `json:"title"`
}

func (job JobReq) Validate() error {
	return validation.ValidateStruct(&job,
		validation.Field(&job.Title, validation.NotNil),
	)
}

type User struct {
	Name  *string `json:"name" `
	Email *string `json:"email" `
	Age   *int    `json:"age,omitempty" `
	Job   *JobReq `json:"job"`
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
		validation.Field(&u.Name),
		validation.Field(&u.Email),
		validation.Field(&u.Age, validation.NotNil, validation.Max(100)),
		validation.Field(&u.Job, validation.NotNil),
	)
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	AddSource: true,
}))

type JSONSerializer struct {
	echo.DefaultJSONSerializer
}

func (d JSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(i)
	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
	} else if se, ok := err.(*json.SyntaxError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
	}

	if err != nil {
		return err
	}

	if v, ok := i.(Validator); ok {
		err := v.Validate()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	return nil
}

func main() {
	e := echo.New()
	e.JSONSerializer = &JSONSerializer{}

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

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err != nil {
				// Define a custom error response
				var code = http.StatusInternalServerError
				var message interface{} = "Internal Server Error"

				// If the error is an echo.HTTPError, use the error code and message
				if he, ok := err.(*echo.HTTPError); ok {
					code = he.Code
					message = he.Message
				}

				// Send the custom error response
				return c.JSON(code, map[string]interface{}{
					"error":   true,
					"message": message,
				})
			}
			return nil
		}
	})

	e.POST("/person", func(c echo.Context) (err error) {
		u := &User{}
		if err := c.Bind(u); err != nil {
			return err
		}

		u_dto := &UserDto{}
		copier.Copy(u_dto, u)
		logger.Info("User: ", "u_dto", u_dto)

		return c.JSON(http.StatusOK, u)
	},
	)

	e.Logger.Fatal(e.Start(":1323"))
}
