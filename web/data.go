package web

import (
	"archive/zip"
	"net/http"
	"path"

	"github.com/go-chi/chi"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
)

func MountData(r chi.Router, rc *zip.ReadCloser) {
	zfs := zipfs.New(rc, "facebook-json.zip")
	hfs := http.FileServer(httpfs.New(zfs))
	r.Mount("/data", http.StripPrefix("/data", hfs))
}

func dataPath(name string) string {
	return path.Join("/data", name)
}
