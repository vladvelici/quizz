package app

import (
	"net/http"
)

// Account GET controller.
// Must enforce MustAuth in routes.
func AccountGet(c *C) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.ViewParams["PageTitle"] = "Quizz - My account"
		c.RenderHTML(w, http.StatusOK, "account")
	})
}

// Account POST controller.
// Must enforce MustAuth in routes.
func AccountPost(c *C) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Validate the new DisplayName, and update the user model.
		c.Ok(w, r)
	})
}
