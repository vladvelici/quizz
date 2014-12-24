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
func MustAuth(f Ctrl) Ctrl {
	ret := func(w http.ResponseWriter, r *http.Request, ctx *C) {
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
		ctx.CurrentUser, err = UserFetchOrCreate(c, usr)
		if err != nil {
			c.Criticalf("Error fetching/creating user: %s", err)
			RenderError(w, http.StatusInternalServerError)
			return
		}
		// MustAuth return url after logout is the homepage
		logoutUrl, _ := user.LogoutURL(c, "/")
		ctx.PageParam("LogoutURL", logoutUrl)
		ctx.PageParam("User", ctx.CurrentUser)
		f(w, r, ctx)
	}

	return Ctrl(ret)
}

// If user exists in the Appengine context, add it to ctx.
// If there's no user, just keep going...
func Auth(f Ctrl) Ctrl {
	ret := func(w http.ResponseWriter, r *http.Request, ctx *C) {
		c := appengine.NewContext(r)
		usr := user.Current(c)
		returnUrl := r.URL.RequestURI()
		if usr != nil {
			var err error
			ctx.CurrentUser, err = UserFetchOrCreate(c, usr)
			if err != nil {
				c.Criticalf("Error fetching/creating user: %s", err)
				RenderError(w, http.StatusInternalServerError)
				return
			}
			logoutUrl, _ := user.LogoutURL(c, returnUrl)
			ctx.PageParam("LogoutURL", logoutUrl)
			ctx.PageParam("User", ctx.CurrentUser)
		}
		loginUrl, _ := user.LoginURL(c, returnUrl)
		ctx.PageParam("LoginURL", loginUrl)
		f(w, r, ctx)
	}

	return Ctrl(ret)
}
