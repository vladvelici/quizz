package app

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handle)
	http.Handle("/usr/", Mid(handle2, MustAuth))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handle 1 updated")
}

func handle2(c *C) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello dear user. %#v", *c.CurrentUser)
	})
}
