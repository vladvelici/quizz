package app

import (
	"gopkg.in/unrolled/render.v1"

	"appengine"
)

var rndr *render.Render

func init() {
	RegisterRoutes()
	rndr = render.New(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".html"},
		IndentJSON: appengine.IsDevAppServer(),
	})
}
