package web

import (
	"archive/zip"
	"net/http"
	"path"

	"github.com/go-chi/chi"
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
)

func MountData(r chi.Router, rcs ...*zip.ReadCloser) {
	ns := make(vfs.NameSpace)
	for _, rc := range rcs {
		zfs := zipfs.New(rc, "facebook-json.zip")
		ns.Bind("/data", zfs, "/", vfs.BindAfter)
	}

	r.Mount("/data", http.FileServer(httpfs.New(ns)))
}

func dataPath(name string) string {
	return path.Join("/data", name)
}
