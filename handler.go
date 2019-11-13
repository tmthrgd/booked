package main

import (
	"archive/zip"
	"net/http"

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
		case isMessageJSON(f):
			if err := parseMessageJSON(f); err != nil {
				return nil, err
			}
		}
	}

	sortMessages()

	fs := zipfs.New(rc, file)
	return http.FileServer(httpfs.New(fs)), nil
}
