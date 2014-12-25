package app

import (
	"github.com/vladvelici/quizz/infra"
	"net/http"
)

// HomePage controller
func HomePage(w http.ResponseWriter, r *http.Request, c *infra.C) {
	c.RenderHTML(w, http.StatusOK, "homepage")
}
