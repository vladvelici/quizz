package app

import (
	"fmt"
	"net/http"
)

func init() {
	RegisterRoutes()
}

func HomePage(c *C) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the home page. The user if any: %#v", c.CurrentUser)
	})
}
