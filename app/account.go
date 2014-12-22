package app

import (
	"fmt"
	"net/http"
)

// Account GET controller.
// Must enforce MustAuth in routes.
func AccountGet(c *C) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello dear user. %#v", *c.CurrentUser)
	})
}

// Account POST controller.
// Must enforce MustAuth in routes.
func AccountPost(c *C) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello dear user. This is a POST method. %#v", *c.CurrentUser)
	})
}
