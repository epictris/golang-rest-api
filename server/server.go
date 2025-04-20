package server

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/go-playground/validator/v10"
	"tris.sh/go/api/errors"
	"tris.sh/go/api/routes"
)

func jsonValidator[A any, R any](
	session *sql.DB, endpoint func(*sql.DB, A) (R, error),
) func(http.ResponseWriter, *http.Request) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return func(w http.ResponseWriter, r *http.Request) error {
		var args A
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(body, &args); err != nil {
			return errors.APIErrorBadRequest("Invalid arguments")
		}
		if err := validate.Struct(args); err != nil {
			return errors.APIErrorBadRequest("Invalid arguments")
		}
		response, err := endpoint(session, args)
		if err != nil {
			return err
		}

		result, err := json.Marshal(response)
		if err != nil {
			return err
		}
		if err := validate.Struct(response); err != nil {
			return err
		}

		if _, err := w.Write(result); err != nil {
			return err
		}

		return nil
	}
}

func errorHandler(endpoint func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := endpoint(w, r); err != nil {
			log.Println(err)
			var error errors.APIError
			if re, ok := err.(errors.APIError); ok {
				error = re
			} else {
				error = errors.APIErrorInternalServerError()
			}
			w.WriteHeader(error.StatusCode())
			if error_response, err := json.Marshal(error); err == nil {
				w.Write(error_response)
			}
		}
	})
}

func registerRoute[A any, R any](
	mux *http.ServeMux,
	db *sql.DB,
	pattern string,
	endpoint func(session *sql.DB, args A) (R, error),
) {
	handler := otelhttp.WithRouteTag(pattern, errorHandler(jsonValidator(db, endpoint)))
	mux.Handle(pattern, handler)
}

func NewHTTPHandler(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	registerRoute(mux, db, "/api/create_user", routes.CreateUser)

	// add HTTP logging instrumentation for all routes
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}
