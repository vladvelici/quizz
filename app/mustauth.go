package app

import (
	"net/http"

	"appengine"
	"appengine/user"
)

// Middleware function to force authentication.
//
// If the user is not authenticated, it redirects to the login page,
// which after sucessfull authentication returns the user to the initial
// page requested.
//
// This function does not check for existing users in the datastore.
func MustAuth(ctx *C, f http.Handler) http.Handler {
	ret := func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		usr := user.Current(c)
		if usr == nil {
			// redirect to authentication
			returnUrl := r.URL.RequestURI()
			loginUrl, _ := user.LoginURL(c, returnUrl)
			w.Header().Set("Location", loginUrl)
			w.WriteHeader(http.StatusFound)
			return
		}

		var err error
		ctx.CurrentUser, err = FetchOrCreate(c, usr)
		if err != nil {
			c.Criticalf("Error fetching/creating user: %s", err)
			RenderError(w, http.StatusInternalServerError)
			return
		}

		f.ServeHTTP(w, r)
	}

	return http.HandlerFunc(ret)
}

// If user exists in the Appengine context, add it to ctx.
// If there's no user, just keep going...
func Auth(ctx *C, f http.Handler) http.Handler {
	ret := func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		usr := user.Current(c)

		var err error
		ctx.CurrentUser, err = FetchOrCreate(c, usr)
		if err != nil {
			c.Criticalf("Error fetching/creating user: %s", err)
			RenderError(w, http.StatusInternalServerError)
			return
		}

		f.ServeHTTP(w, r)
	}

	return http.HandlerFunc(ret)
}
