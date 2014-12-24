package app

import (
	"net/http"
)

// Account GET controller.
// Must enforce MustAuth in routes.
func AccountGet(w http.ResponseWriter, r *http.Request, c *C) {
	c.PageParam("PageTitle", "Quizz - My account")
	c.Params["User"] = c.CurrentUser
	c.Render(w, http.StatusOK, "account")
}

// Account POST controller.
// Must enforce MustAuth in routes.
func AccountPost(w http.ResponseWriter, r *http.Request, c *C) {
	// TODO: Validate the new DisplayName, and update the user model.
	c.Ok(w, r)
}
