// +build !dev

package web

import (
	"html/template"

	"github.com/shurcooL/httpfs/html/vfstemplate"
	"go.tmthrgd.dev/booked/internal/templates"
)

func newTemplate(name string, funcs template.FuncMap) Template {
	t := template.New(name).Funcs(templateFuncs).Funcs(funcs)
	return template.Must(vfstemplate.ParseFiles(templates.FileSystem, t, name))
}
