package main

import (
	"archive/zip"
	"net/http"

	"go.tmthrgd.dev/booked/messages"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
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

	fs := zipfs.New(rc, file)
	return http.FileServer(httpfs.New(fs)), nil
}
