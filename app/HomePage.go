package app

import (
	"net/http"
)

// HomePage controller
func HomePage(w http.ResponseWriter, r *http.Request, c *C) {
	c.RenderHTML(w, http.StatusOK, "homepage")
}
