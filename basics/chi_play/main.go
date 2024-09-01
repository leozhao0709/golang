package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/render"
	"github.com/google/go-cmp/cmp"
	"github.com/jellydator/validation"
)

type Validator interface {
	Validate() error
}

type JobReq struct {
	Title *string `json:"title"`
}

func (job JobReq) Validate() error {
	return validation.ValidateStruct(&job,
		validation.Field(&job.Title, validation.NotNil),
	)
}

func (job JobReq) String() string {
	return fmt.Sprintf("Title: %s", *job.Title)
}

type PersonReq struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
	Age   *int    `json:"age"`
	Job   *JobReq `json:"job"`
}

func (u PersonReq) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.NotNil),
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Age, validation.Min(13)),
		validation.Field(&u.Job, validation.NotNil),
	)
}

func ValidateReqBody[v Validator]() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			validator := new(v)

			if err := render.Decode(r, validator); err != nil {
				render.JSON(w, r, render.M{
					"error": err.Error(),
				})
				return
			}

			ctx := context.WithValue(r.Context(), "req_body", validator)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	logger := httplog.NewLogger("httplog-example", httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		SourceFieldName: "source",
	})
	r.Use(httplog.RequestLogger(logger))

	// r.Use(middleware.Logger)

	render.Decode = func(r *http.Request, v interface{}) error {
		var err error

		switch render.GetRequestContentType(r) {
		case render.ContentTypeJSON:
			err = DecodeJSON(r.Body, v)
		default:
			err = render.DefaultDecoder(r, v)
		}

		if err != nil {
			return err
		}

		if validator, ok := v.(Validator); ok {
			return validator.Validate()
		}

		return nil
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "Hello World!")
	})

	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("./static"))
		http.StripPrefix("/static", fs).ServeHTTP(w, r)
	})

	r.
		With(ValidateReqBody[PersonReq]()).
		Post("/person", func(w http.ResponseWriter, r *http.Request) {
			p := r.Context().Value("req_body").(*PersonReq)
			s := []*PersonReq{p}
			logger := httplog.LogEntry(r.Context())

			s_c := cmp.Diff(nil, s)
			logger.Info("s list", "s_c", s_c)

			render.JSON(w, r, s)
		})
	http.ListenAndServe(":3000", r)
}

// DecodeJSON decodes a given reader into an interface using the json decoder.
func DecodeJSON(r io.Reader, v interface{}) error {
	defer io.Copy(io.Discard, r) //nolint:errcheck
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}
