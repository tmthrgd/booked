package main

import (
	"archive/zip"
	"net/http"

	"github.com/go-chi/chi"
	"go.tmthrgd.dev/booked/messages"
	"go.tmthrgd.dev/booked/web"
)

func open(file string) (http.Handler, error) {
	rc, err := zip.OpenReader(file)
	if err != nil {
		return nil, err
	}

	for _, f := range rc.File {
		switch {
		case messages.IsMessageJSON(f):
			if err := messages.ParseMessageJSON(f); err != nil {
				return nil, err
			}
		}
	}

	r := chi.NewRouter()
	r.NotFound(web.NotFoundHandler())
	r.MethodNotAllowed(web.MethodNotAllowedHandler())

	web.MountData(r, rc)

	return r, nil
}
