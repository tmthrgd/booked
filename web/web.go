package web

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"html/template"
	"io"
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

var errorTmpl = NewTemplate("error.tmpl", template.FuncMap{
	"httpStatusText": http.StatusText,
})

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
		errorMsg := html.EscapeString(err.Error())
		switch {
		case errors.Is(err, os.ErrNotExist):
			statusCode = http.StatusNotFound
			errorMsg = "The requested file was not found."
		case err == errMethodNotAllowed:
			statusCode = http.StatusMethodNotAllowed
			errorMsg = fmt.Sprintf("The request method <code>%s</code> is inappropriate for the URL <code>%s</code>.",
				html.EscapeString(r.Method), html.EscapeString(r.URL.Path))
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(statusCode)

		var buf bytes.Buffer
		if errorTmpl.Execute(&buf, &struct {
			StatusCode int
			Message    template.HTML
		}{
			statusCode,
			template.HTML(errorMsg),
		}) == nil {
			buf.WriteTo(w)
			return
		}

		fmt.Fprintf(w, "%d %s: %s", statusCode,
			http.StatusText(statusCode), errorMsg)
	}
}

type Template interface {
	Execute(io.Writer, interface{}) error
}

func NewTemplate(name string, funcs template.FuncMap) Template {
	return newTemplate(name, funcs)
}

var templateFuncs = template.FuncMap{
	"assetPath": assetPath,
	"dataPath":  dataPath,
}
