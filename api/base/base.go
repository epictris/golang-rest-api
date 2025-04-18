package base

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"tris.sh/go/api/errors"
)

func jsonValidator[A any, R any](
	session *sql.DB, endpoint func(*sql.DB, A) (R, error),
) func(http.ResponseWriter, *http.Request) error {
	validate := validator.New()
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

func errorHandler(endpoint func(http.ResponseWriter, *http.Request) error) http.Handler {
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

func RegisterRoute[A any, R any](
	path string,
	session *sql.DB,
	endpoint func(*sql.DB, A) (R, error),
) {
	http.Handle(path, errorHandler(jsonValidator(session, endpoint)))
}
