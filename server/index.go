package server

import (
	"html/template"
	"net/http"

	"github.com/peterhellberg/generator-generator/generator"
)

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func indexHandler(ctx *Context, r *http.Request, w http.ResponseWriter) error {
	var gs []generator.Generator

	sess := ctx.Session.Clone()
	defer sess.Close()

	c := sess.DB("").C("generators")

	err := c.Find(nil).All(&gs)
	if err != nil {
		return err
	}

	return index.Execute(w,
		struct{ Generators []generator.Generator }{gs})
}