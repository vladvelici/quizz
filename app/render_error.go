package app

import (
	"fmt"
	"net/http"
)

func RenderError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Ooops. Something went wrong.")
}
