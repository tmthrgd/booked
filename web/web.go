package web

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// NotFoundHandler returns a handler that serves a 404 Not Found error.
func NotFoundHandler() http.HandlerFunc {
	return ErrorHandler(func(http.ResponseWriter, *http.Request) error {
		return os.ErrNotExist
	})
}

var errMethodNotAllowed = errors.New("the request method is inappropriate for the requested URL")

// MethodNotAllowedHandler returns a handler that serves a 405 Method Not
// Allowed error.
func MethodNotAllowedHandler() http.HandlerFunc {
	return ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		return errMethodNotAllowed
	})
}

// ErrorHandler converts a handler with an error return to a http.HandlerFunc,
// sending a HTTP error code and a JSON formatted RFC 7807 problem detail
// document to the client appropriate for a given error.
func ErrorHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		statusCode := http.StatusInternalServerError
		errorMsg := err.Error()
		switch {
		case errors.Is(err, os.ErrNotExist):
			statusCode = http.StatusNotFound
			errorMsg = "The requested file was not found."
		case err == errMethodNotAllowed:
			statusCode = http.StatusMethodNotAllowed
			errorMsg = fmt.Sprintf("The request method %s is inappropriate for the URL %s.",
				r.Method, r.URL.Path)
		}

		http.Error(w, errorMsg, statusCode)
	}
}
