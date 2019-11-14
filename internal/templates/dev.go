// +build dev

package templates

import "net/http"

// FileSystem contains project templates.
var FileSystem http.FileSystem = http.Dir("templates")
