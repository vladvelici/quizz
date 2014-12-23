package app

import (
	"net/http"
)

// HomePage controller
func HomePage(c *C) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.RenderHTML(w, http.StatusOK, "homepage")
	})
}
