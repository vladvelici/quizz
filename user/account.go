package user

import (
	"github.com/vladvelici/quizz/infra"

	"net/http"
)

// Account GET controller.
// Must enforce MustAuth in routes.
func AccountGet(w http.ResponseWriter, r *http.Request, c *infra.C) error {
	c.PageParam("PageTitle", "Quizz - My account")
	c.Params["User"] = c.CurrentUser
	c.Render(w, http.StatusOK, "account")
	return nil
}

// Account POST controller.
// Must enforce MustAuth in routes.
func AccountPost(w http.ResponseWriter, r *http.Request, c *infra.C) error {
	// TODO: Validate the new DisplayName, and update the user model.
	c.Ok(w, r)
	return nil
}
