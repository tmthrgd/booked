package main

import (
	"archive/zip"
	"net/http"

	"github.com/go-chi/chi"
	"go.tmthrgd.dev/booked/messages"
	"go.tmthrgd.dev/booked/web"
)

func open(files ...string) (http.Handler, error) {
	var rcs []*zip.ReadCloser
	for _, file := range files {
		rc, err := zip.OpenReader(file)
		if err != nil {
			return nil, err
		}
		rcs = append(rcs, rc)

		for _, f := range rc.File {
			switch {
			case messages.IsMessageJSON(f):
				if err := messages.ParseMessageJSON(f); err != nil {
					return nil, err
				}
			}
		}
	}

	r := chi.NewRouter()
	r.NotFound(web.NotFoundHandler())
	r.MethodNotAllowed(web.MethodNotAllowedHandler())

	web.MountAssets(r)
	web.MountData(r, rcs...)

	messages.Mount(r)

	if err := messages.FixupURIs(rcs...); err != nil {
		return nil, err
	}

	// TODO(tmthrgd): Index page.
	r.Get("/", http.RedirectHandler("/messages/", http.StatusTemporaryRedirect).ServeHTTP)

	return r, nil
}
