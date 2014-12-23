package app

import (
	"net/http"
)

// Middleware type
type Middleware func(*C, http.Handler) http.Handler

// Controller creator type
type Ctrl func(*C) http.Handler

// Request context
type C struct {
	CurrentUser *User
	ViewParams  map[string]interface{}
}

func (c *C) RenderHTML(w http.ResponseWriter, status int, template string) {
	rndr.HTML(w, status, template, c.ViewParams)
}

func (c *C) RenderHTMLData(w http.ResponseWriter, status int, template string, data interface{}) {
	rndr.HTML(w, status, template, data)
}

func (c *C) RenderJSON(w http.ResponseWriter, status int, data interface{}) {
	rndr.JSON(w, status, data)
}

// Redirect or output JSON "status: ok".
func (c *C) Ok(w http.ResponseWriter, r *http.Request) {
	if IsApi(r) {
		c.RenderJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	} else {
		// TODO: add success flash message
		c.Redirect(w, r.URL.RequestURI())
	}
}

// Redirect or output JSON "status: fail".
func (c *C) Fail(w http.ResponseWriter, r *http.Request) {
	if IsApi(r) {
		c.RenderJSON(w, http.StatusOK, map[string]string{"status": "fail"})
	} else {
		// TODO: add failure flash message
		c.Redirect(w, r.URL.RequestURI())
	}
}

// Redirect or output given JSON data.
func (c *C) Data(w http.ResponseWriter, r *http.Request, data interface{}) {
	if IsApi(r) {
		c.RenderJSON(w, http.StatusOK, data)
	} else {
		// TODO: add success flash message
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
	ret := func(w http.ResponseWriter, r *http.Request) {
		ctx := new(C)
		ctx.ViewParams = make(map[string]interface{})
		ctx.ViewParams["PageTitle"] = "Quizz"
		f := h(ctx)
		for i := len(mid) - 1; i >= 0; i-- {
			f = mid[i](ctx, f)
		}
		f.ServeHTTP(w, r)
	}
	return http.HandlerFunc(ret)
}
