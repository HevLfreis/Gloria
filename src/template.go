package main

import (
	"github.com/flosch/pongo2"
	"net/http"
)

var busTpl = pongo2.Must(pongo2.FromFile("templates/bus.html"))

func Render(w http.ResponseWriter, tpl *pongo2.Template, context map[string]interface{}) error {
	SetResponse(w, "html")
	context = pongo2.Context(context)
	err := tpl.ExecuteWriter(context, w)
	return err
}

