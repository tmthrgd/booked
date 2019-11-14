// +build dev

package web

import (
	"html/template"
	"io"

	"github.com/shurcooL/httpfs/html/vfstemplate"
	"go.tmthrgd.dev/booked/internal/templates"
)

func newTemplate(name string, funcs template.FuncMap) Template {
	return &devTemplate{name, funcs}
}

type devTemplate struct {
	name  string
	funcs template.FuncMap
}

func (dt *devTemplate) Execute(w io.Writer, data interface{}) error {
	t := template.New(dt.name).Funcs(templateFuncs).Funcs(dt.funcs)

	if _, err := vfstemplate.ParseFiles(templates.FileSystem, t, dt.name); err != nil {
		return err
	}

	return t.Execute(w, data)
}
