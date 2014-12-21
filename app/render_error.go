package app

import (
	"fmt"
	"net/http"
)

func RenderInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "There was an internal server error. Please try again.")
}

func RenderError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Ooops. Something went wrong.")
}
