package web

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/shurcooL/httpfs/filter"
	handlers "github.com/tmthrgd/httphandlers"
	"go.tmthrgd.dev/booked/internal/assets"
	"go.tmthrgd.dev/vfshash"
)

var assetNames = vfshash.NewAssetNames(assets.FileSystem)

func MountAssets(r chi.Router) {
	assetNamesH := http.FileServer(assetNames).ServeHTTP
	r.With(
		handlers.SetHeadersWrap(map[string]string{
			"Content-Type":  "image/png",
			"Cache-Control": "public, max-age=1209600", // 14 days
		}),
	).Get("/favicon.ico", assetNamesH)

	if assetNames.IsContentAddressable() {
		r = r.With(
			handlers.NeverModified,
			handlers.SetHeaderWrap("Cache-Control", "public, max-age=31536000, immutable"), // 1 year
		)
	} else {
		r = r.With(handlers.SetHeaderWrap("Cache-Control", "public, no-cache"))
	}

	fs := http.FileServer(filter.Skip(assets.FileSystem, excludeAssets))
	r.Mount("/assets", http.StripPrefix("/assets", fs))
}

func excludeAssets(path string, info os.FileInfo) bool {
	return info.IsDir() || strings.HasPrefix(info.Name(), ".")
}
