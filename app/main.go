package app

import (
	"github.com/gorilla/sessions"
	"gopkg.in/unrolled/render.v1"

	"appengine"
)

var rndr *render.Render
var store = sessions.NewCookieStore([]byte(CookieSecret))

func init() {
	RegisterRoutes()

	rndr = render.New(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".html"},
		IndentJSON: appengine.IsDevAppServer(),
	})
}
