// +build dev

package assets

import "net/http"

// FileSystem contains project assets.
var FileSystem http.FileSystem = http.Dir("assets")
