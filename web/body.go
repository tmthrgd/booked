package web

import (
	"bytes"
	"net/http"
	"strconv"
)

func WriteResponseBody(w http.ResponseWriter, b []byte) error {
	hdr := w.Header()
	hdr.Set("Content-Length", strconv.Itoa(len(b)))

	if n, err := w.Write(b); n == 0 && err != nil {
		// Only forward the error if nothing was written.
		hdr.Del("Content-Length")
		return err
	}

	return nil
}

func WriteTemplateResponse(w http.ResponseWriter, tmpl Template, data interface{}) error {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return WriteResponseBody(w, buf.Bytes())
}
