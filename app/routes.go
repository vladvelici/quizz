package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes() {
	r := mux.NewRouter().StrictSlash(false)

	r.Handle("/", Mid(HomePage, Auth))

	account := r.Path("/account").Subrouter()
	account.Methods("GET").Handler(Mid(AccountGet, MustAuth))
	account.Methods("POST").Handler(Mid(AccountPost, MustAuth))

	http.Handle("/", r)
}
