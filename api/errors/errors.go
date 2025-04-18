package errors

import "net/http"

type APIError interface {
	Error() string
	StatusCode() int
}

type apiError struct {
	Message    string `json:"error"`
	statusCode int
}

func (m apiError) Error() string {
	return m.Message
}

func (m apiError) StatusCode() int {
	return m.statusCode
}

func APIErrorBadRequest(message string) APIError {
	return apiError{message, http.StatusBadRequest}
}

func APIErrorUnauthorized() APIError {
	return apiError{"Unauthorized", http.StatusUnauthorized}
}

func APIErrorForbidden(message string) APIError {
	return apiError{message, http.StatusForbidden}
}

func APIErrorNotFound(message string) APIError {
	return apiError{message, http.StatusNotFound}
}

func APIErrorInternalServerError() APIError {
	return apiError{"Internal Server Error", http.StatusInternalServerError}
}
