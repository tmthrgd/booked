// +build generate

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"
	"go.tmthrgd.dev/booked/internal/assets"
	"go.tmthrgd.dev/vfshash"
)

func main() {
	if err := vfsgen.Generate(vfshash.NewFileSystem(assets.FileSystem), vfsgen.Options{
		Filename:        "internal/assets/vfsdata.go",
		PackageName:     "assets",
		BuildTags:       "!dev",
		VariableName:    "FileSystem",
		VariableComment: "FileSystem contains project assets.",
	}); err != nil {
		log.Fatal(err)
	}
}
