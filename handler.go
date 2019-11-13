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

	fs := zipfs.New(rc, file)
	return http.FileServer(httpfs.New(fs)), nil
}
