package app

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Middleware type
type Middleware func(Ctrl) Ctrl

// Controller creator type
type Ctrl func(w http.ResponseWriter, r *http.Request, c *C)

// Convenience method to return a http.HandlerFunc using a context.
func (f Ctrl) HandleFunc(c *C) http.HandlerFunc {
	h := func(w http.ResponseWriter, r *http.Request) {
		f(w, r, c)
	}
	return http.HandlerFunc(h)
}

// Request context
type C struct {
	CurrentUser *User
	Params      map[string]interface{}
	IsApi       bool
	Session     *sessions.Session
}

// Add a parameter to c.Params only if it's an api call.
func (c *C) ApiParam(key string, value interface{}) {
	if c.IsApi == true {
		c.Params[key] = value
	}
}

// Add a parameter to c.Params only if it's a page call.
func (c *C) PageParam(key string, value interface{}) {
	if c.IsApi == false {
		c.Params[key] = value
	}
}

// Render JSON if c.IsApi == true, or HTML otherwise. Uses c.Params as data.
//
// This method is suitable for GET requests, where some data needs to be given
// back on both API and website requests.
//
// For POST requests, use c.Ok, c.Fail, or c.Data, as they render JSON for
// API calls and add a flash message and redirect for website requests.
func (c *C) Render(w http.ResponseWriter, status int, template string) {
	if c.IsApi {
		c.RenderJSON(w, status)
	} else {
		c.RenderHTML(w, status, template)
	}
}

// Same as c.Render(), except it uses the given data instead of c.Params.
func (c *C) RenderData(w http.ResponseWriter, r *http.Request, status int, template string, data interface{}) {
	if c.IsApi {
		c.RenderJSONData(w, status, data)
	} else {
		c.RenderHTMLData(w, status, template, data)
	}
}

// Render HTML using c.Params as data.
func (c *C) RenderHTML(w http.ResponseWriter, status int, template string) {
	rndr.HTML(w, status, template, c.Params)
}

// Render HTML using the given data.
func (c *C) RenderHTMLData(w http.ResponseWriter, status int, template string, data interface{}) {
	rndr.HTML(w, status, template, data)
}

// Render JSON using c.Params.
func (c *C) RenderJSON(w http.ResponseWriter, status int) {
	rndr.JSON(w, status, c.Params)
}

// Render JSON using given data.
func (c *C) RenderJSONData(w http.ResponseWriter, status int, data interface{}) {
	rndr.JSON(w, status, data)
}

// Redirect or output JSON "status: ok".
func (c *C) Ok(w http.ResponseWriter, r *http.Request) {
	if IsApi(r) {
		c.RenderJSONData(w, http.StatusOK, map[string]string{"status": "ok"})
	} else {
		c.Session.AddFlash("All done!", "_flash_success")
		c.Redirect(w, r.URL.RequestURI())
	}
}

// Redirect or output JSON "status: fail".
func (c *C) Fail(w http.ResponseWriter, r *http.Request) {
	if IsApi(r) {
		c.RenderJSONData(w, http.StatusOK, map[string]string{"status": "fail"})
	} else {
		c.Session.AddFlash("Something went wrong :(", "_flash_failure")
		c.Redirect(w, r.URL.RequestURI())
	}
}

// Redirect or output given JSON data.
func (c *C) Data(w http.ResponseWriter, r *http.Request, data interface{}) {
	if IsApi(r) {
		c.RenderJSONData(w, http.StatusOK, data)
	} else {
		c.Session.AddFlash(data, "_flash_success")
		c.Redirect(w, r.URL.RequestURI())
	}
}

// Use this function to redirect after POST. It sets the header Location: to.
// HTTP Status used is 303 See Other.
func (c *C) Redirect(w http.ResponseWriter, to string) {
	w.Header().Set("Location", to)
	w.WriteHeader(http.StatusSeeOther)
}

// If the request has the header "Accept: application/json" or the query
// string parameter "format=json", this function returns true.
func IsApi(r *http.Request) bool {
	if r.Header.Get("Accept") == "application/json" {
		return true
	}

	query := r.URL.Query()
	if format, ok := query["format"]; ok {
		if len(format) == 1 && format[0] == "json" {
			return true
		}
	}

	return false
}

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
