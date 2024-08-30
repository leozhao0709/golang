package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jellydator/validation"
)

type JobReq struct {
	Title *string `json:"title"`
}

func (job JobReq) Validate() error {
	return validation.ValidateStruct(&job,
		validation.Field(&job.Title, validation.NotNil),
	)
}

// func (job JobReq) Bind(r *http.Request) error {
// 	return validation.ValidateStruct(&job,
// 		validation.Field(&job.Title, validation.NotNil),
// 	)
// }

type PersonReq struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
	Age   *int    `json:"age"`
	Job   *JobReq `json:"job"`
}

func (u PersonReq) Bind(r *http.Request) error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.NotNil),
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Age, validation.Min(13)),
		validation.Field(&u.Job, validation.NotNil),
	)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	render.Decode = func(r *http.Request, v interface{}) error {
		var err error

		switch render.GetRequestContentType(r) {
		case render.ContentTypeJSON:
			err = DecodeJSON(r.Body, v)
		case render.ContentTypeXML:
			err = render.DecodeXML(r.Body, v)
		case render.ContentTypeForm:
			err = render.DecodeForm(r.Body, v)
		default:
			err = errors.New("render: unable to automatically decode the request content type")
		}

		return err
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.
		// With(DisallowUnknownJsonFields[PersonReq]()).
		Post("/person", func(w http.ResponseWriter, r *http.Request) {
			// p := r.Context().Value("request_body").(*PersonReq)

			p := &PersonReq{}
			if err := render.Bind(r, p); err != nil {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, map[string]string{"error": err.Error()})
				return
			}

			render.JSON(w, r, p)
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
