package infra

import (
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/unrolled/render.v1"

	"appengine"
)

// Renderer
var rndr = render.New(render.Options{
	Directory:  "templates",
	Layout:     "layout",
	Extensions: []string{".html"},
	IndentJSON: appengine.IsDevAppServer(),
})

// Cookie store
var store = sessions.NewCookieStore([]byte(CookieSecret))

// Controller type
type Ctrl func(w http.ResponseWriter, r *http.Request, c *C) error

// Convenience method to return a http.HandlerFunc using a context.
func (f Ctrl) HandlerFunc(c *C) http.HandlerFunc {
	h := func(w http.ResponseWriter, r *http.Request) {
		// todo: deal with error.
		f(w, r, c)
	}
	return http.HandlerFunc(h)
}

// Middleware type
type Middleware func(Ctrl) Ctrl

// Create a context, create the controller using the context, and
// create a http.Handler that combines all the middleware.
func Mid(h Ctrl, mid ...Middleware) http.Handler {
	f := h
	for i := len(mid) - 1; i >= 0; i-- {
		f = mid[i](f)
	}
	ret := func(w http.ResponseWriter, r *http.Request) {
		ctx := new(C)
		ctx.IsApi = IsApi(r)
		ctx.Params = make(map[string]interface{})
		ctx.PageParam("PageTitle", "Quizz")
		ctx.Session, _ = store.Get(r, "main-session")
		f(w, r, ctx)
	}
	return http.HandlerFunc(ret)
}
